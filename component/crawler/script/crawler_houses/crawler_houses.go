package main

import (
	"bigCitySmallHouse/component/crawler"
	"bigCitySmallHouse/component/crawler/factory"
	"bigCitySmallHouse/model/house"
	"bigCitySmallHouse/mongodb"
	"bigCitySmallHouse/mongodb/collection"
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
	houses, err := fetchAll()
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

func fetchAll() ([]house.House, error) {
	log.Println("获取所有的房源数据...")

	houseMap := make(map[string]house.House) // 这里用map是为了去除重复数据，乐有家爬取到的数据有一小部分重复
	var allHouses []house.House

	page := 0
	for {
		page++
		tHouses, info, err := fetchPage(page)
		if err != nil {
			return nil, err
		}
		log.Printf("fetchPage %d/%d\n", page, info.TotalPage)
		allHouses = append(allHouses, tHouses...)
		houses2map(houseMap, tHouses)
		if page >= info.TotalPage || info.IsLastPage {
			break
		}
	}

	log.Println("总抓取数据长度（可能包含重复）", len(allHouses))
	houses := houseMap2Arr(houseMap)

	return houses, nil
}

func houses2map(houseMap map[string]house.House, houses []house.House) {
	for _, h := range houses {
		_, exist := houseMap[h.SourceId]
		if exist {
			log.Printf("存在重复source_id: %s\n", h.SourceId)
			log.Printf("被覆盖的详细数据: %v\n", h)
		}
		houseMap[h.SourceId] = h
	}
}

func houseMap2Arr(houseMap map[string]house.House) []house.House {
	houses := make([]house.House, 0, len(houseMap))
	for _, h := range houseMap {
		houses = append(houses, h)
	}
	return houses
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
