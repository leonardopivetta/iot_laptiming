package laptiming

import (
	"context"
	"io/ioutil"
	"laptiming/backend/firebase_interface"
	"laptiming/backend/telegram_interface"
	"log"
	"math/rand"
	"strconv"

	"github.com/gin-gonic/gin"
	proto "github.com/golang/protobuf/proto"
	"google.golang.org/api/iterator"
)

var firebaseApp *firebase_interface.FirebaseApp;

func publisLapTimeToFirestore(carId int32, lapTime int64, timestamp int64){
	firebaseApp.GetLapTimesCollection().Add(context.Background(), map[string]interface{}{
		"id": carId ,
		"lapTime": lapTime,
		"timestamp": timestamp,
	});
}

func publishPitStatusToFirestore(carId int32, pitStatus PitStatus, timestamp int64){
	firebaseApp.GetPitStatusCollection().Add(context.Background(), map[string]interface{}{
		"id": carId ,
		"pitStatus": pitStatus,
		"timestamp": timestamp,
	});
}

func setupLapTimeRoutes(router *gin.Engine){
	lapTimeGroup := router.Group("/laptime");
	// firestoreClient := firebaseApp.GetFirestoreClient();
	lapTimeGroup.GET("/", func(ctx *gin.Context) { //TODO: change to POST
		car := rand.Intn(10);
		//TODO: add protobuf body parsing
		telegram_interface.NotifyCar(car, "Laptime: 1:23.456");
		firebaseApp.GetLapTimesCollection().Add(ctx, map[string]interface{}{
			"id": car ,
			"lapTime": rand.Intn(1000),
			"timestamp": rand.Intn(1000000),
		});
		ctx.String(200, "pushed: "+strconv.Itoa(car) +" to firestore");
	});
	lapTimeGroup.POST("/", func(ctx *gin.Context){
		protomessage := LapTimeMessage{};
		body, err := ioutil.ReadAll(ctx.Request.Body);
		if err != nil {
			log.Print("reading body error: ", err)
			ctx.Status(400);
			return;
		}
		err = proto.Unmarshal(body, &protomessage);
		if err != nil {
			log.Print("Received: ", body);

			log.Print("unmarshaling error: ", err)
			ctx.Status(400);
			return;
		}
		log.Print("Received: ", protomessage);
		go publisLapTimeToFirestore(protomessage.GetId(), protomessage.GetLapTime(), protomessage.GetTimestamp());
		go telegram_interface.NotifyCar(int(protomessage.GetId()), "Laptime: "+strconv.FormatInt(protomessage.GetLapTime(), 10));
		ctx.Status(200);
	});
	lapTimeGroup.GET("/last", func(ctx *gin.Context){
			doc,_ := firebaseApp.GetLapTimesCollection().OrderBy("timestamp", firebase_interface.Desc).Limit(1).Documents(ctx).Next();
			ctx.JSON(200, doc.Data());
	})

	lapTimeGroup.GET("/car", func(ctx *gin.Context){
		carId := ctx.Query("id");
		if carId == ""{
			ctx.String(400, "Missing id parameter");
			return
		}
		//parse carId to int
		car, err := strconv.ParseInt(carId, 10, 64);
		if err != nil{
			ctx.String(400, "Invalid id parameter");
			return;
		}
		iter := firebaseApp.GetLapTimesCollection().Where("id", "==", car).Documents(ctx);
		docs := make([]map[string]interface{}, 0);
		for {
			doc, err := iter.Next();
			if err == iterator.Done {
                break
        	}
			if err != nil {
				log.Fatalf("Failed to iterate: %v", err);
			}
			docs = append(docs, doc.Data());
		}
		ctx.JSON(200, docs);
	});	
}

func setupPitStatusRoutes(router *gin.Engine){
	pitStatusGroup := router.Group("/pitstatus");
	pitStatusGroup.GET("/", func(ctx *gin.Context){
		ctx.String(200, "Pit status");
	});
	pitStatusGroup.POST("/", func(ctx *gin.Context){
		protomessage := PitStatusMessage{};
		body, err := ioutil.ReadAll(ctx.Request.Body);
		if err != nil {
			log.Print("Received: ", body);
			log.Print("reading body error: ", err)
			ctx.Status(400);
			return;
		}
		err = proto.Unmarshal(body, &protomessage);
		if err != nil {
			log.Print("unmarshaling error: ", err)
			ctx.Status(500);
			return;
		}
		log.Print("Received: ", protomessage);

		go publishPitStatusToFirestore(protomessage.GetId(), protomessage.GetPitStatus(), protomessage.GetTimestamp());
		go telegram_interface.NotifyCar(int(protomessage.GetId()), "Pit status: "+ PitStatus_name[int32(protomessage.GetPitStatus())]);	
		ctx.Status(200);
	});
}

func init(){
	firebaseApp = &firebase_interface.FirebaseApp{};
	firebaseApp.Init();
}

func SetupRouter(router *gin.Engine){
	setupLapTimeRoutes(router);
	setupPitStatusRoutes(router);
}
