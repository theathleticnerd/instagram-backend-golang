package DBSetup

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitConnection() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://job:y0SFqFVTmz9uDBEY@cluster0.xpnzm.mongodb.net/instaDB?retryWrites=true&w=majority"))
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to Database")
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	return client
}
