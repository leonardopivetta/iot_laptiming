package telegram_interface

import (
	"fmt"
	"log"
	"strings"
)

func execCommand(message TelegramMessage){

	// split the message into command and content
	splitted := strings.Split(message.Message.Text, " ");
	command := splitted[0];
	var content string;
	if len(splitted) > 1 {
		content = strings.Join(splitted[1:], " ");
	}

	// execute the command
	var responseText string;
	switch command {
	case "/listen":
		responseText = addListenerToCar(content, message.Message.Chat.Id);
		break;
	case "/showListeners":
		responseText = getAllListeners();
		break;
	default:
		responseText = "Unknown command";
	}
	resMessage := TelegramSendMessage{
	ChatId: message.Message.Chat.Id,
	Text: responseText,
	}
	go sendMessage(&resMessage);
}

var CarListeners = make(map[int](map[int]bool));

func addListenerToCar(content string, user int) string {
	//parse the content as an int
	var carId int;
	_, err := fmt.Sscanf(content, "%d", &carId);
	if err != nil {
		log.Println(err);
		return "not a number";
	}
	if(CarListeners[carId] == nil) {
		CarListeners[carId] = make(map[int]bool);
	}
	CarListeners[carId][user] = true;
	return "Done";
}

func getAllListenersForCar(carId int) string {
	var response string;
	for user, _ := range CarListeners[carId] {
		response += fmt.Sprintf("- %d\n", user);
	}
	return response;
}

func getAllListeners() string {
	var response string = "";
	//iterate through all cars
	for carId, _  := range CarListeners {
		response += fmt.Sprintf("Car %d:\n", carId);
		response += getAllListenersForCar(carId);
		response += "\n";
	}
	if response == "" {
		response = "No listeners";
	}
	return response;
}

func NotifyCar(carId int, message string) {
	for user, _ := range CarListeners[carId] {
		resMessage := TelegramSendMessage{
			ChatId: user,
			Text: message,
		}
		go sendMessage(&resMessage);
	}
}