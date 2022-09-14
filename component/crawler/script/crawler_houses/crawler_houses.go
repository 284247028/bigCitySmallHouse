package main

import (
	"bigCitySmallHouse/component/crawler"
	"bigCitySmallHouse/component/crawler/factory"
	"bigCitySmallHouse/component/crawler/script"
	"bigCitySmallHouse/model/house"
	"bigCitySmallHouse/mongodb"
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
	res, err := script.HousesInsert(houses)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res)
}

func fetchAll() ([]house.House, error) {
	log.Println("fetchAll...")

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
