package main

import (
	house2 "bigCitySmallHouse/component/crawler/model/house"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MongoClient *mongo.Client

var House1 = house2.House{
	SourceId: "111",
	Source:   house2.SourceAnjuke,
	Type:     house2.TypeApartment,
	Name:     "house1",
}

var House2 = house2.House{
	SourceId: "222",
	Source:   house2.SourceLeyoujia,
	Type:     house2.TypeResidence,
	Name:     "house2",
}

var House3 = house2.House{
	SourceId: "333",
	Source:   house2.SourceAnjuke,
	Type:     house2.TypeVilla,
	Name:     "house3",
}

var House4 = house2.House{
	SourceId: "444",
	Source:   house2.SourceLeyoujia,
	Type:     house2.TypeApartment,
	Name:     "house4",
}

var House5 = house2.House{
	SourceId: "555",
	Source:   house2.SourceLeyoujia,
	Type:     house2.TypeApartment,
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
	qHouse := house2.House{
		SourceId: "dddd",
	}
	one, err := MongoClient.Database("crawler").Collection("house").InsertOne(context.TODO(), qHouse)
	if err != nil {
		return err
	}
	log.Println(one)
	return nil
}
