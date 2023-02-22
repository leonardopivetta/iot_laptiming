#pragma once


#include <ESP8266WiFi.h>
#include <ESP8266HTTPClient.h>
#include <WiFiClient.h>
#include <string.h>

#include "../nanopb/pb_common.h"
#include "../nanopb/pb_decode.h"
#include "../nanopb/pb_encode.h"
#include "../proto/messages.pb.h"

#include "queue.h"
typedef struct ToBackendMessage {
    const char* endpoint;
    const char* data;
} ToBackendMessage_t;


class BackendConnector {

    protected:
        String serverName = "http://192.168.43.126:8079";

        bool ensureConnection();

        HTTPClient client;
        WiFiClient wifiClient;

        Queue<ToBackendMessage_t*> messageQueue;

        void sendMessage(ToBackendMessage_t *message);

        void emptyQueue();
    public:
         const char* ssid = "Note 9 Leo"; 
         const char* password = "leonardo";

        BackendConnector();
        BackendConnector(const char* ssid, const char* password, String serverName);

        void sendLaptime(iot_laptiming_LapTimeMessage *lapTime);
        void sendPitStatus(iot_laptiming_PitStatusMessage *pitStatus);
    
};



#define SERIALIZE_TO_STRING(in, type, message) \
    char* s_out = in; \
    uint8_t buffer[type##_size];\
    pb_ostream_t stream = pb_ostream_from_buffer(buffer, type##_size);\
    pb_encode(&stream, type##_fields, message);\
    for (size_t i = 0; i < stream.bytes_written; i++) {\
        *s_out = buffer[i];\
        s_out++;\
    }\
    *s_out = '\0';

