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
	houses, err := load()
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

func load() ([]house.House, error) {
	log.Println("从数据库读取需要填充的房源...")
	opts := collection.NewOptions()
	opts.DB = mongodb.DBCrawler
	opts.Collection = house.CollectionName
	tCollection := collection.NewCollection(opts)
	filter := bson.D{}
	if _source != "" {
		filter = append(filter, bson.E{"source", _source})
	}

	findOpts := options.Find().SetProjection(bson.D{
		{"_id", 1},
		{"source_id", 1},
		{"uid", 1},
		{"source", 1},
	})
	cursor, err := tCollection.MCollection().Find(context.TODO(), filter, findOpts)
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
	log.Println("填充房源详情...")
	param := crawler.NewSingleParam()
	tHouses := make([]house.House, 0, len(houses))
	count := 0
	total := len(houses)
	for _, h := range houses {
		count++
		param.Id = h.SourceId
		source := h.Source
		if _source != "" {
			source = house.Source(_source)
		}
		houseFactory, err := factory.NewFactory(source)
		if err != nil {
			return nil, err
		}
		parser := houseFactory.CreateSingleParser(param)
		tHouse, err := parser.Parse()
		if err != nil {
			log.Printf("抓取详情失败: %s %d/%d - %s\n", err.Error(), count, total, h.Id.Hex())
			continue
		}
		log.Printf("抓取详情成功 %d/%d - %s\n", count, total, tHouse.UId)
		tHouse.Id = h.Id
		tHouses = append(tHouses, *tHouse)
	}
	return tHouses, nil
}
