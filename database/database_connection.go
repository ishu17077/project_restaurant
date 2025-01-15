package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client{
	mongoDb := "mongodb://localhost:27017" 
	fmt.Print(mongoDb)
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second) //? 10 seconds
	defer cancel()
	client,err :=mongo.Connect(ctx, options.Client().ApplyURI(mongoDb))
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to mongoDb")
	return client
}

var Client *mongo.Client = DBInstance();

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection{
	var collection *mongo.Collection = client.Database("restaurant").Collection(collectionName)
	return collection
}