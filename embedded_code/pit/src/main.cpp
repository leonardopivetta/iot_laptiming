#include <Arduino.h>
#include <IRremote.h>
#include <SoftwareSerial.h>


// #define DEBUG_FLAG
#define ENTRY_FLAG
// #define EXIT_FLAG

//! REMEMBER TO CHANGE THE platformio.ini FILE TO MATCH THE PIT YOU ARE COMPILING FOR

#if defined(ENTRY_FLAG) && defined(EXIT_FLAG)
#error "You can't define both ENTRY_FLAG and EXIT_FLAG"
#endif


//! ARDUINO LEONARDO (pit entry)
#ifdef ENTRY_FLAG
#define IR_RECEIVE_PIN 8
#define PIN_SEND 7
#define CODE 0xe1      //Code used to send to the car the entrance in the pit lane
#endif

//! ARDUINO PRO MICRO (pit exit)
#ifdef EXIT_FLAG
#define IR_RECEIVE_PIN 7                              
#define PIN_SEND 4
#define CODE 0x8e     //Code used to send to the car the exit from the pit lane
#endif

// The time that the car should not be in the field of view after the receive in [[ms]]
#define IR_OVERLAY 200 

// Limit of the number of cars
#define MAX_CARS 10

uint32_t previous_send = 0;

int receiveIr();
void sendSerial(int car);
void sendIr();


void setup() {
  // ------------------ IR CONFIGURATION ------------------
  IrSender.begin(PIN_SEND);
  IrReceiver.begin(IR_RECEIVE_PIN);

  // ------------------ SERIAL CONFIGURATION ------------------
  Serial1.begin(9600, SERIAL_8N1);
}

void loop() {
  int car = receiveIr();
  if(car > -1){
    sendSerial(car);
  }
  sendIr();
}


// Reads the IR signal and returns the car index
int receiveIr(){
  if (IrReceiver.decode()) {
    IrReceiver.resume();
    uint16_t carIndex = IrReceiver.decodedIRData.command; 
    if(carIndex < 1 || carIndex > MAX_CARS + 1 || carIndex == CODE){  // -1 because if the ir reads nothing returns 0
      return -1;
    }
    carIndex-=1;  // Array displacement for memory usage
    return carIndex;
  }
  return -1;
}

// Sends the car index to the start/finish (esp8266) line
void sendSerial(int car){
  Serial1.println(car);
}

// Sends the IR signal to the car to change the status
void sendIr(){
  if(millis()-previous_send > 500){
    IrReceiver.disableIRIn();
    IrSender.sendNEC(CODE, CODE, 2);
    IrReceiver.enableIRIn();
    previous_send= millis();
  }
}