package tests

import (
	"bigCitySmallHouse/component/crawler/model/house"
	"bigCitySmallHouse/mongodb"
	"bigCitySmallHouse/mongodb/collection"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"testing"
)

func TestUpsert(t *testing.T) {
	err := mongodb.NewDB().ConnectMongodb(mongodb.NewOptions())
	if err != nil {
		t.Fatal(err)
	}
	coll := collection.NewCollection(collection.NewOptions())

	houses := make([]house.House, 0)
	houses = append(houses, house.House{
		SourceId: "111222333",
	})

	houses = append(houses, house.House{
		SourceId: "222333444",
	})

	for _, h := range houses {

		marshal, err := bson.Marshal(h)
		if err != nil {
			t.Fatal(err)
			return
		}

		m := bson.M{}
		err = bson.Unmarshal(marshal, &m)
		if err != nil {
			t.Fatal(err)
		}

		update := bson.M{
			"$set": h,
		}
		//opts := &options.ReplaceOptions{}
		//opts.SetUpsert(true)
		replaceOne, err := coll.MCollection().UpdateOne(context.TODO(), bson.D{{"source_id", h.SourceId}}, update)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(replaceOne)
	}
	log.Println("end")
}
