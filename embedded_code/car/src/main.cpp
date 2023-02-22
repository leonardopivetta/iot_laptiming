#include <Arduino.h>
#include <IRremote.h>

// #define DEBUG_FLAG

#define CAR_0
// #define CAR_2  

//! REMEMBER TO CHANGE THE platformio.ini FILE TO MATCH THE CAR YOU ARE COMPILING FOR

#if defined(CAR_0) && defined(CAR_2)
#error "You can't define both CAR_0 and CAR_2"
#endif

#ifdef CAR_0

// ! CAR 0 (arduino nano)
#define IR_RECEIVE_PIN 5
#define PIN_SEND 7

#define RED_PIN 2
#define GREEN_PIN 3
#define BLUE_PIN 4     

#define CAR_ID 1

#endif

#ifdef CAR_2

// ! CAR 2 (arduino pro micro)
#define IR_RECEIVE_PIN 7
#define WHITE_LED_PIN 6                                 
#define PIN_SEND 4
#define CAR_ID 3

#define RED_PIN 5
#define GREEN_PIN 8
#define BLUE_PIN 9

#endif     


uint32_t previous_send = 0;

enum car_status {PIT, OUT_LAP, LAP};

car_status currentStatus = PIT;

void receiveIr();
void sendIr();
void writeStatusLed();



void setup() {

  // ------------------ IR CONFIGURATION ------------------
  IrSender.begin(PIN_SEND);
  IrReceiver.begin(IR_RECEIVE_PIN);

  // ------------------ SERIAL CONFIGURATION ------------------
  Serial.begin(9600);
  
  // ------------------ RGB LED CONFIGURATION ------------------
  pinMode(RED_PIN,   OUTPUT);
  pinMode(GREEN_PIN, OUTPUT);
  pinMode(BLUE_PIN,  OUTPUT);
}

unsigned long prevWrite = 0;

void loop() {
  receiveIr();
  sendIr();
  unsigned long currentMillis = millis();
  if(currentMillis - prevWrite > 200){
    prevWrite = currentMillis;
    writeStatusLed();
  }
}


// Updates the current status based on the IR signal received
void receiveIr(){
  if (IrReceiver.decode()) {
    IrReceiver.resume();
    switch (IrReceiver.decodedIRData.command)
    {
    case 0xfa:
      currentStatus = LAP;
      #ifdef DEBUG_FLAG
      Serial.print("la macchina ");
      Serial.print(CAR_ID);
      Serial.println(" è passata dal via! ;P");
      #endif
      break;
    case 0xe1:
      currentStatus = PIT;
      #ifdef DEBUG_FLAG
      Serial.print("la macchina ");
      Serial.print(CAR_ID);
      Serial.println(" è entrata nei pit! B)");
      #endif
      break;
    case 0x8e:
      currentStatus = OUT_LAP;
      #ifdef DEBUG_FLAG
      Serial.print("la macchina ");
      Serial.print(CAR_ID);
      Serial.println(" è uscita dai pit! :D");
      #endif
      break;
    default:
      break;
    }
  }
}

// Notifies the own id
void sendIr(){
  if(millis()-previous_send > 500){
    IrReceiver.disableIRIn();
    IrSender.sendNEC(CAR_ID, CAR_ID, 2);
    IrReceiver.enableIRIn();
    previous_send= millis();
  }
}

// Writes the current status on the led
void writeStatusLed(){
  switch (currentStatus){
  case PIT:
    //Setting the led RED for PIT
    digitalWrite(RED_PIN,   LOW);
    digitalWrite(GREEN_PIN, HIGH);
    digitalWrite(BLUE_PIN,  HIGH);
    break;
  case OUT_LAP:
    //Setting the led GREEN for OUT_LAP
    digitalWrite(RED_PIN,   HIGH);
    digitalWrite(GREEN_PIN, LOW);
    digitalWrite(BLUE_PIN,  HIGH);
    break;
  case LAP:
    //Setting the led BLUE for LAP
    digitalWrite(RED_PIN,   HIGH);
    digitalWrite(GREEN_PIN, HIGH);
    digitalWrite(BLUE_PIN,  LOW);
    break;
  }
}