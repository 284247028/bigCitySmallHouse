package main

import (
	"bigCitySmallHouse/model/house"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)

func TestInsertOne(t *testing.T) {
	err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	result, err := MongoClient.Database("crawler").Collection("house").InsertOne(context.TODO(), House1)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(result)
}

func TestInsertMany(t *testing.T) {
	err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	houses := []interface{}{House2, House3}
	result, err := MongoClient.Database("crawler").Collection("house").InsertMany(context.TODO(), houses)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(result)
}

func TestUpdateOne(t *testing.T) {
	err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	filter := bson.D{
		{"source_id", "111"},
	}
	house1 := House1
	house1.Name = "house111"
	update := bson.D{
		{"$set", house1},
	}
	result, err := MongoClient.Database("crawler").Collection("house").UpdateOne(context.TODO(), filter, update)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(result)
}

func TestUpsertOne(t *testing.T) {
	err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	filter := bson.D{
		{"source_id", "444"},
	}
	update := bson.D{
		{"$set", House4},
	}
	opts := &options.UpdateOptions{}
	opts.SetUpsert(true)
	result, err := MongoClient.Database("crawler").Collection("house").UpdateOne(context.TODO(), filter, update, opts)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(result)
}

func TestUpdateById(t *testing.T) {
	err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	house4 := House4
	house4.Name = "house444"
	update := bson.D{
		{"$set", house4},
	}
	idStr := "63254aab02be1cfea788c41b"
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return
	}
	result, err := MongoClient.Database("crawler").Collection("house").UpdateByID(context.TODO(), id, update)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(result)
}

func TestUpsertById(t *testing.T) {
	err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	update := bson.D{
		{"$set", House5},
	}
	opts := &options.UpdateOptions{}
	opts.SetUpsert(true)
	id := primitive.NewObjectID()
	result, err := MongoClient.Database("crawler").Collection("house").UpdateByID(context.TODO(), id, update, opts)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(result)
}

func TestUpdateMany(t *testing.T) {
	err := Connect()
	if err != nil {
		t.Fatal(err)
	}
	filter := bson.D{
		{"source", house.SourceLeyoujia},
	}
	update := bson.D{
		{"$set", bson.D{
			{"type", house.TypeVilla},
		}},
	}
	result, err := MongoClient.Database("crawler").Collection("house").UpdateMany(context.TODO(), filter, update)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(result)
}
