package main

import (
	"bigCitySmallHouse/component/crawler"
	"bigCitySmallHouse/component/crawler/factory"
	"bigCitySmallHouse/model/house"
	"bigCitySmallHouse/mongodb"
	"bigCitySmallHouse/mongodb/collection"
	"bigCitySmallHouse/mongodb/collections"
	"context"
	"flag"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var _source string

func init() {
	flag.StringVar(&_source, "source", "", "数据来源")
	flag.Parse()
	opts := mongodb.NewOptions()
	err := mongodb.NewDB().ConnectMongodb(opts)
	if err != nil {
		return
	}
}

func main() {
	houses, err := fetch()
	if err != nil {
		log.Fatalln(err)
	}
	houses, err = fillHouses(houses)
	if err != nil {
		log.Fatalln(err)
	}
	opts := &collection.Options{}
	opts.DB = mongodb.DBCrawler
	opts.Collection = house.CollectionName
	houseCollection := collections.NewCollectionHouse(opts)
	results, err := houseCollection.HouseUpsertMany(houses)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(results)
}

func fetch() ([]house.House, error) {
	log.Println("fetching...")
	opts := collection.NewOptions()
	opts.DB = mongodb.DBCrawler
	opts.Collection = house.CollectionName
	tCollection := collection.NewCollection(opts)
	findOpts := options.Find().SetProjection(bson.D{
		{"_id", 1},
		{"source_id", 1},
	})
	cursor, err := tCollection.MCollection().Find(context.TODO(), bson.D{}, findOpts)
	if err != nil {
		return nil, err
	}

	var houses []house.House
	err = cursor.All(context.TODO(), &houses)
	if err != nil {
		return nil, err
	}
	return houses, nil
}

func fillHouses(houses []house.House) ([]house.House, error) {
	log.Println("filling...")
	houseFactory, err := factory.NewFactory(house.Source(_source))
	if err != nil {
		return nil, err
	}
	param := crawler.NewSingleParam()
	tHouses := make([]house.House, 0, len(houses))
	count := 0
	total := len(houses)
	for _, h := range houses {
		count++
		log.Printf("filling %d/%d - %s\n", count, total, h.SourceId)
		param.Id = h.SourceId
		parser := houseFactory.CreateSingleParser(param)
		tHouse, err := parser.Parse()
		if err != nil {
			log.Println(err)
			continue
		}
		tHouse.Id = h.Id
		tHouses = append(tHouses, *tHouse)
	}
	return tHouses, nil
}
