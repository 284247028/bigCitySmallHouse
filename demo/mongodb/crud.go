package main

import (
	"bigCitySmallHouse/model/house"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var mongoClient *mongo.Client

func main() {
	err := connect()
	if err != nil {
		log.Fatalln(err)
	}

	err = insert()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("success")

}

func connect() error {
	uri := "mongodb://localhost:27017"
	clientOptions := options.Client()
	clientOptions.ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	mongoClient = client
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}
	return nil
}

func insert() error {
	qHouse := house.House{
		SourceId: "dddd",
	}
	one, err := mongoClient.Database("crawler").Collection("house").InsertOne(context.TODO(), qHouse)
	if err != nil {
		return err
	}
	log.Println(one)
	return nil
}
