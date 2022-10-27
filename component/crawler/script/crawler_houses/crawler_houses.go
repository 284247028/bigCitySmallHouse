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
	err := mongodb.NewDB().ConnectMongodb(opts)
	if err != nil {
		return
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

		packs := Houses2Packs(tHouses)
		_, err = houseCollection.PackUpsertMany(packs)
		if err != nil {
			log.Printf("保存第%d页数据出错，错误信息：%s\n", page, err.Error())
		}

		log.Printf("保存成功\n")
		count += len(tHouses)
	}

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

func Houses2Packs(houses []house.House) []house.Pack {
	packs := make([]house.Pack, 0, len(houses))
	for _, tHouse := range houses {
		pack := house.Pack{
			Status: house.PackStatusList,
			House:  tHouse,
		}
		packs = append(packs, pack)
	}
	return packs
}
