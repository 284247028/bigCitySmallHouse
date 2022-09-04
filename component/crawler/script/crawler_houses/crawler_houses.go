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
	param := crawler.ListParam{
		Page: 1,
	}
	houseFactory, err := factory.NewFactory(house.Source(_source))
	if err != nil {
		return nil, err
	}
	parser := houseFactory.CreateListParser(&param)
	info, err := parser.Info()
	if err != nil {
		return nil, err
	}

	houseMap := make(map[string]house.House) // 这里用map是为了去除重复数据，乐有家爬取到的数据有一小部分重复

	for i := 1; i <= info.TotalPage; i++ {
		log.Printf("fetchPage %d/%d\n", i, info.TotalPage)
		tHouses, err := fetchPage(i)
		if err != nil {
			return nil, err
		}
		houses2map(houseMap, tHouses)
	}

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

func fetchPage(page int) ([]house.House, error) {
	houseFactory, err := factory.NewFactory(house.Source(_source))
	if err != nil {
		return nil, err
	}
	param := crawler.ListParam{
		Page: page,
	}
	parser := houseFactory.CreateListParser(&param)
	return parser.Parse()
}
