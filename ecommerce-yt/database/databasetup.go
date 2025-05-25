package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func DBSet() *mongo.Client {
	// Set up MongoDB connection
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}


	// Check the connection
	// ctx,cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// 	defer cancel()
	// 	err = client.Connect(ctx)
	// if err != nil {
	// 		log.Fatal(err)
	// 	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("faild to Connected to MongoDB! :(")
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	// Set up the user collection
	// var collection *mongo.Collection = client.Database("Ecommerce").Collection(collectionName)  //old style
	collection := client.Database("Ecommerce").Collection(collectionName)
	return collection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	// Set up the product collection
	collection := client.Database("Ecommerce").Collection(collectionName)
	return collection
}
