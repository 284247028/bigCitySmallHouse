package main

import (
	"bigCitySmallHouse/component/crawler"
	"bigCitySmallHouse/component/crawler/factory"
	"bigCitySmallHouse/component/crawler/model/house"
	"bigCitySmallHouse/mongodb"
	"bigCitySmallHouse/mongodb/collections"
	"context"
	"flag"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

var _source string

func init() {
	flag.StringVar(&_source, "source", "", "数据来源")
	flag.Parse()
	opts := mongodb.NewOptions()
	opts.Uri = "mongodb://admin:admin@43.138.174.42:27017/"
	err := mongodb.NewDB().ConnectMongodb(opts)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	packs, err := load()
	if err != nil {
		log.Fatalln(err)
	}
	packs, err = fillHouses(packs)
	if err != nil {
		log.Fatalln(err)
	}
	houseCollection := collections.NewCollectionPack(nil)
	results, err := houseCollection.PackUpsertMany(packs)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(results)
}

func load() ([]house.Pack, error) {
	log.Println("从数据库读取需要填充的房源...")
	tCollection := collections.NewCollectionPack(nil)
	filter := bson.D{}
	if _source != "" {
		filter = append(filter, bson.E{"house.source", _source})
	}

	//findOpts := options.Find().SetProjection(bson.D{
	//	{"_id", 1},
	//	{"source_id", 1},
	//	{"uid", 1},
	//	{"source", 1},
	//})
	cursor, err := tCollection.MCollection().Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	var packs []house.Pack
	err = cursor.All(context.TODO(), &packs)
	if err != nil {
		return nil, err
	}
	return packs, nil
}

func fillHouses(packs []house.Pack) ([]house.Pack, error) {
	log.Println("填充房源详情...")
	param := crawler.NewSingleParam()
	tPacks := make([]house.Pack, 0, len(packs))
	count := 0
	total := len(packs)
	for _, p := range packs {
		count++
		param.Id = p.House.SourceId
		source := p.House.Source
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
			log.Printf("抓取详情失败: %s %d/%d - %s\n", err.Error(), count, total, p.House.UId)
			continue
		}
		log.Printf("抓取详情成功 %d/%d - %s\n", count, total, tHouse.UId)
		//tHouse.Id = h.Id
		p.House = *tHouse
		p.Status = house.PackStatusSingle
		tPacks = append(tPacks, p)
	}
	return tPacks, nil
}
