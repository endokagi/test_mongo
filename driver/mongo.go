package driver

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb+srv://endozaki:thewlip123456@cluster0.rp9vy.mongodb.net/sample_mflix?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	collection := client.Database("test").Collection("users")

	return collection
}
