package main

import (
	"bigCitySmallHouse/component/crawler"
	"bigCitySmallHouse/component/crawler/factory"
	"bigCitySmallHouse/component/crawler/model/house"
	"bigCitySmallHouse/mongodb"
	"bigCitySmallHouse/mongodb/collections"
	"flag"
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
	fetchAll()
}

func fetchAll() {
	log.Printf("获取房源 %s 数据...\n", _source)

	houseCollection := collections.NewCollectionPack(nil)

	page := 0
	count := 0
	for {
		page++
		tHouses, info, err := fetchPage(page)
		if err != nil {
			log.Printf("爬取第%d页出错，错误信息：%s\n", page, err.Error())
			continue
		}
		log.Printf("爬取成功 %d/%d\n", page, info.TotalPage)
		if page >= info.TotalPage || info.IsLastPage {
			break
		}

		packs := houseCollection.Houses2Packs(tHouses)
		_, err = houseCollection.PackUpsertMany(packs)
		if err != nil {
			log.Printf("保存第%d页数据出错，错误信息：%s\n", page, err.Error())
		}

		log.Printf("保存成功\n")
		count += len(tHouses)
	}
	// 实际数量低于总数，因为第三方房源有id重复，被覆盖，正常
	log.Println("爬取完成，总数：", count)
}

func fetchPage(page int) ([]house.House, *crawler.ListInfo, error) {
	houseFactory, err := factory.NewFactory(house.Source(_source))
	if err != nil {
		return nil, nil, err
	}
	param := crawler.ListParam{
		Page: page,
	}
	parser := houseFactory.CreateListParser(&param)
	return parser.Parse()
}
