package firebase_interface

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go"
	"google.golang.org/api/option"
)

type FirebaseApp struct {
	app *firebase.App;
}

func (firebaseApp *FirebaseApp) Init(){
	opt := option.WithCredentialsFile("./secrets/iot-laptiming-firebase-adminsdk-xtxis-7b0046ec1a.json");
	fmt.Printf("opt: %v\n", opt);
	app, err := firebase.NewApp(context.Background(), nil, opt);
	if err != nil {
		log.Fatalf("Error initializing firebase app: %s", err);
	}
	firebaseApp.app = app;
}

func (firebaseApp *FirebaseApp) GetFirestoreClient() *firestore.Client{
	client, err := firebaseApp.app.Firestore(context.Background());
	if err != nil {
		log.Fatalf("Error getting Firestore client: %s", err);
	}
	return client;
}

func (firebaseApp *FirebaseApp) GetLapTimesCollection() *firestore.CollectionRef{
	client := firebaseApp.GetFirestoreClient();
	return client.Collection("lapTimes");
}

func (firebaseApp *FirebaseApp) GetPitStatusCollection() *firestore.CollectionRef{
	client := firebaseApp.GetFirestoreClient();
	return client.Collection("pitStatus");
}

const Desc = firestore.Desc;
