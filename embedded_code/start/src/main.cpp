#include <Arduino.h>
#include <IRremote.h>
#include <SoftwareSerial.h>

#include "./http/backendConnector.hpp"

//Flag used for degub portion of the code
// #define DEBUG_FLAG

#define PIN_SEND D2
#define IR_RECEIVE_PIN D3
#define CODE 0xfa

#define SS_RX D6
#define SS_TX D5

SoftwareSerial pitEntrySerial;
SoftwareSerial pitExitSerial;

// The time that the car should not be in the field of view after the receive in [[ms]]
#define IR_OVERLAY 200 

// Limit of the number of cars
#define MAX_CARS 10

enum car_status {
  PIT,
  OUT_LAP,
  LAP
}; 

typedef struct {
  uint16_t start;
  uint16_t lap_timing;
  car_status status;
} carData_t;

carData_t carDatas[MAX_CARS];

uint16_t previous_send = 0;

BackendConnector* backendConnector;

void receiveIR();
void sendSerial(int car);
void sendIR();

#define PIT(x, y, status)  \
  String y##read = x.readStringUntil('\n');\
  if(y##read.length()){\
    char* detectedCar = (char*) y##read.c_str();\
    while(detectedCar != NULL && *detectedCar != '\0' && *detectedCar != '\n'){\
      if(*detectedCar < '0' || *detectedCar > '9'){detectedCar++; continue;}\
      carDatas[*detectedCar - '0'].status = status;\
      iot_laptiming_PitStatusMessage message;\
      message.id = *detectedCar - '0';\
      message.pitStatus = iot_laptiming_PitStatus_##y;\
      message.timestamp = millis();\
      if(false){\
      backendConnector->sendPitStatus(&message);}\
      detectedCar++;\
    }\
  }

void setup()  
{ 
  // --------------------------  IR INITIALIZATION  --------------------------
  IrSender.begin(PIN_SEND);
  IrReceiver.begin(IR_RECEIVE_PIN);
  // -------------------------- SERIAL INITIALIZATION --------------------------
  Serial.begin(9600);
  pitEntrySerial.begin(9600, SWSERIAL_8N1, SS_RX, SS_TX, false);
  pitExitSerial.begin(9600, SWSERIAL_8N1, D7, -1, false);
  #ifdef DEBUG_FLAG
  if(!pitEntrySerial){
    Serial.println("Failed to start SoftwareSerial pitEntrySerial");
    while(1){delay(1000);}
  }
  if(!pitExitSerial){
    Serial.println("Failed to start SoftwareSerial pitExitSerial");
    while(1){delay(1000);}
  }
  #endif

  // -------------------------- CAR DATA INITIALIZATION --------------------------
  const uint16_t currentMillis = millis();
  for(unsigned int i = 0; i < MAX_CARS; i++){
    carDatas[i] = {
      currentMillis,
      0
    };
  }

  // -------------------------- WIFI INITIALIZATION --------------------------
  backendConnector = new BackendConnector();
  WiFi.begin(backendConnector->ssid, backendConnector->password);
  #ifdef DEBUG_FLAG
  Serial.print("Connecting to WiFi");
  int i = 0;
  #endif
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    #ifdef DEBUG_FLAG
    Serial.print(".");
    if((i++) % 16 == 0){
      Serial.println();
    }
    #endif
  }
  #ifdef DEBUG_FLAG
  Serial.println("Connected to the WiFi network in " + String(i*500) + "ms");
  Serial.println();
  #endif
}  
                 
void loop()  
{  
  receiveIR();
  sendIr();
  PIT(pitExitSerial, PIT_EXIT, OUT_LAP)
  PIT(pitEntrySerial, PIT_ENTRY, PIT)
}



void receiveIR(){
  if (IrReceiver.decode()) {
    IrReceiver.resume(); //resets data previously memorized from the receiver
    uint16_t carIndex = IrReceiver.decodedIRData.command; 
    if(carIndex < 1 || carIndex > MAX_CARS + 1 || carIndex == CODE){  // -1 because if the ir reads nothing returns 0
      return;
    }
    carIndex-=1;  // Array displacement for memory usage
    carData_t* current_car = &carDatas[carIndex];
    unsigned long currentMillis = millis(); 
    if(currentMillis-current_car->start > IR_OVERLAY){ // Check if the car time should be updated or if it's too soon
      current_car->lap_timing = currentMillis - current_car->start;
      current_car->start = currentMillis;
      
      #ifdef DEBUG_FLAG
      Serial.print("Lap time:  \t");
      Serial.print(current_car->lap_timing);
      Serial.print("\tCAR: ");
      Serial.println(carIndex);
      #endif

      switch(current_car->status){
        case OUT_LAP:
          // Set the car status to lap
          current_car->status = LAP;
          break;
        case LAP:
          // Send the lap time to the backend
          iot_laptiming_LapTimeMessage message;
          message.id = carIndex;
          message.lapTime = current_car->lap_timing;
          message.timestamp = millis();
          backendConnector->sendLaptime(&message);
          break;
      }
    }
  }
}

void sendIr(){
  if(millis()-previous_send > 500){
    IrReceiver.disableIRIn();
    IrSender.sendNEC(CODE, CODE, 2);
    IrReceiver.enableIRIn();
    previous_send = millis();
  }  
}