package telegram_interface

import (
	"bytes"
	"context"
	"encoding/json"

	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"

	"github.com/gin-gonic/gin"

	"io/ioutil"
)

var currentLink string;
var telegramApiUrl string = "https://api.telegram.org/bot" + getEnv("TELEGRAN_BOT_TOKEN");

func getEnv(key string) string {
	err := godotenv.Load("./secrets/telegram-secrets.env");
	if err != nil {
		log.Println(err)
		log.Fatalf("Error loading .env file");
	}
	return os.Getenv(key);
}

func setupWebhook() error {
	url := telegramApiUrl+"/setWebhook?url="+currentLink+"/telegram/";
	res, err := http.Get(url);
	if err != nil {
		log.Println(err);
		return err;
	}
	if res.StatusCode != 200 {
		log.Println(res.StatusCode);
		return err;
	}
	log.Println("Webhook setup complete");
	return nil;
}

type TelegramSendMessage struct {
	Text string `json:"text"`
	ChatId int `json:"chat_id"`
}

func sendMessage(message *TelegramSendMessage) {
	// encode the message in a json format
	jsonMessage, err := json.Marshal(*message);
	if err != nil {
		log.Println(err);
		return;
	}

	res, err := http.Post("https://api.telegram.org/bot"+getEnv("TELEGRAN_BOT_TOKEN")+"/sendMessage", "application/json", bytes.NewBuffer(jsonMessage));
	if err != nil {
		log.Println(err);
		return;
	}
	if res.StatusCode != 200 {
		body, _ := ioutil.ReadAll(res.Body);
		log.Println(res.StatusCode, string(body));
		return;
	}
}

type TelegramMessage struct {
	UpdateId int `json:"update_id"`
	Message struct {
		MessageId int `json:"message_id"`
		From struct {
			Id uint `json:"id"`
			IsBot bool `json:"is_bot"`
			FirstName string `json:"first_name"`
			LastName string `json:"last_name"`
			Username string `json:"username"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Chat struct {
			Id int `json:"id"`
			FirstName string `json:"first_name"`
			LastName string `json:"last_name"`
			Username string `json:"username"`
			Type string `json:"type"`
		} `json:"chat"`
		Date int `json:"date"`
		Text string `json:"text"`
	} `json:"message"`

}

func Init(ctx context.Context){

	tunnel, err := ngrok.Listen(
		ctx,
		config.HTTPEndpoint(),
		ngrok.WithAuthtoken(getEnv("NGROK_AUTH_TOKEN")),
	)
	if err != nil {
		log.Fatalf("Error creating tunnel: %s", err);
	}
	currentLink = tunnel.URL();
	log.Printf("Tunnel created: %s", currentLink);

	setupWebhook();

	router := gin.Default();
	router.POST("/telegram/", func(ctx *gin.Context) {
		jsonData, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			log.Println(err);
			ctx.Status(400);
			return;
		}
		var message TelegramMessage;
		err = json.Unmarshal(jsonData, &message);
		if err != nil {
			log.Println(err);
			ctx.Status(400);
			return;
		}
		go execCommand(message);
		ctx.Status(200);
	});
	go http.Serve(tunnel, router);
}