#include "./backendConnector.hpp"

bool BackendConnector::ensureConnection() {
    if(WiFi.status() != WL_CONNECTED) {
        WiFi.begin(ssid, password);
        Serial.println("Not connected");
        delay(500);
        return false;
    }
    return true;
}

void BackendConnector::emptyQueue(){
    for(int i = 0; i < messageQueue.count(); i++) {
        ToBackendMessage_t *message = messageQueue.peek();
        messageQueue.pop();
        delete message;
    }
}



BackendConnector::BackendConnector(){
    ensureConnection();
    emptyQueue();
}

BackendConnector::BackendConnector(const char* ssid, const char* password, String serverName){
    if(ssid != NULL){
        this->ssid = ssid;
    }
    if(password != NULL){
        this->password = password;
    }
    if(serverName != NULL){
        this->serverName = serverName;
    }
    BackendConnector();
}

void BackendConnector::sendMessage(ToBackendMessage_t *message){
    messageQueue.push(message);
    if(ensureConnection()){
        const int messageQueueCount = messageQueue.count();
        for(int i = 0; i < messageQueueCount; i++) {
            ToBackendMessage_t *message = messageQueue.peek();
            messageQueue.pop();
            client.begin(wifiClient, serverName+(message->endpoint));
            client.addHeader("Content-Type", "application/protobuf");
            int resCode = client.POST(String(message->data));
            if(resCode != 200 && resCode != 500){
                Serial.print("Error sending message to backend: ");
                Serial.println(resCode);
                messageQueue.push(message);
            }
        }
    }
}


void BackendConnector::sendLaptime(iot_laptiming_LapTimeMessage *lapTime){
    ToBackendMessage_t *message = new ToBackendMessage_t();
    // message->endpoint = (new String("/laptime/"))->c_str();// ! cos'è sta mostruosità?
    message->endpoint = "/laptime/";
    char out[iot_laptiming_LapTimeMessage_size];
    SERIALIZE_TO_STRING(out, iot_laptiming_LapTimeMessage, lapTime)
    message->data = out;
    sendMessage(message);
}

void BackendConnector::sendPitStatus(iot_laptiming_PitStatusMessage *pitStatus){
    ToBackendMessage_t *message = new ToBackendMessage_t();
    message->endpoint = "/pitstatus/";
    char out[iot_laptiming_PitStatusMessage_size];
    SERIALIZE_TO_STRING(out, iot_laptiming_PitStatusMessage, pitStatus)
    message->data = out;
    sendMessage(message);
}