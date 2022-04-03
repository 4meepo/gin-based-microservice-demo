package datasource

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MongoClient *mongo.Client

func init() {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://mongo-user:0mW6JQpR2GPW9fFR@demo.ga4cx.mongodb.net/cooking?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx := context.Background()
	c, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	MongoClient = c
}
