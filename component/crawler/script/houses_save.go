package script

import (
	"bigCitySmallHouse/component/crawler"
	"bigCitySmallHouse/model/house"
	"bigCitySmallHouse/mongodb"
	"bigCitySmallHouse/mongodb/collection"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func HousesInsert(houses []house.House) (*mongo.InsertManyResult, error) {
	log.Println("HousesInsert")
	fHouses := crawler.FormatHouses(houses)
	opts := collection.NewOptions()
	opts.DB = mongodb.DBCrawler
	opts.Collection = house.CollectionName
	tCollection := collection.NewCollection(opts)
	return tCollection.MCollection().InsertMany(context.TODO(), fHouses)
}

func HouseUpdate(houses []house.House) {
	log.Println("HouseUpdate")
	opts := collection.NewOptions()
	opts.DB = mongodb.DBCrawler
	opts.Collection = house.CollectionName
	tCollection := collection.NewCollection(opts)
	for _, h := range houses {
		update := bson.M{
			"$set": h,
		}
		res, err := tCollection.MCollection().UpdateByID(context.TODO(), h.Id, update)
		if err != nil {
			return
		}
		log.Println(res)
	}
}
