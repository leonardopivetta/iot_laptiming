package main

import (
	// "io/ioutil"
	"context"
	"laptiming/backend/telegram_interface"

	// pb "laptiming/backend/laptiming"
	"log"

	//import the package inside the folder iot_laptiming/
	"laptiming/backend/laptiming"

	"github.com/gin-gonic/gin"
	// proto "github.com/golang/protobuf/proto"
	socketio "github.com/googollee/go-socket.io"
)

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

var socketServer *socketio.Server;


func setupSocketIo(){
	socketServer = socketio.NewServer(nil);
	socketServer.OnConnect("/", func(s socketio.Conn) error {
		// s.SetContext("")
		log.Println("connected:", s.ID())
		return nil
	});
	socketServer.OnEvent("/", "update", func(s socketio.Conn, msg string) {
		log.Println("update required");
	});
}




func init(){
	
	socketServer = socketio.NewServer(nil);
}

// func clientNotifier(){
// 	client := firebaseApp.GetFirestoreClient();
// 	defer client.Close();
// 	it := client.Collection("lapTimes").Snapshots(context.Background());
// 	for {
// 		snap, err := it.Next();
// 		if err != nil {
// 			log.Fatalf("Failed to iterate: %v", err);
// 		}
// 		if snap != nil {
// 			socketServer.BroadcastToNamespace("/", "update");
// 		}
// 	}
// }


func main() {
	r := gin.Default();
	r.Use(CORSMiddleware());
	telegram_interface.Init(context.Background());

	// setupLapTime(r);
	setupSocketIo();
	
	go socketServer.Serve();
	defer socketServer.Close();
	// go clientNotifier();
	//bind the socketServer to the router on socket/ path
	r.GET("/socket/*any", gin.WrapH(socketServer))
	laptiming.SetupRouter(r);
	r.Run(":8079");
}
