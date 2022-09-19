package main

import (
	"bigCitySmallHouse/model/house"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MongoClient *mongo.Client

var House1 = house.House{
	SourceId: "111",
	Source:   house.SourceAnjuke,
	Type:     house.TypeApartment,
	Name:     "house1",
}

var House2 = house.House{
	SourceId: "222",
	Source:   house.SourceLeyoujia,
	Type:     house.TypeResidence,
	Name:     "house2",
}

var House3 = house.House{
	SourceId: "333",
	Source:   house.SourceAnjuke,
	Type:     house.TypeVilla,
	Name:     "house3",
}

var House4 = house.House{
	SourceId: "444",
	Source:   house.SourceLeyoujia,
	Type:     house.TypeApartment,
	Name:     "house4",
}

var House5 = house.House{
	SourceId: "555",
	Source:   house.SourceLeyoujia,
	Type:     house.TypeApartment,
	Name:     "house5",
}

func main() {
	err := Connect()
	if err != nil {
		log.Fatalln(err)
	}

	err = insert()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("success")

}

func Connect() error {
	uri := "mongodb://localhost:27017"
	clientOptions := options.Client()
	clientOptions.ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}
	MongoClient = client
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
	one, err := MongoClient.Database("crawler").Collection("house").InsertOne(context.TODO(), qHouse)
	if err != nil {
		return err
	}
	log.Println(one)
	return nil
}
