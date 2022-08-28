package factory

import (
	"bigCitySmallHouse/component/crawler"
	"bigCitySmallHouse/model/house"
	"log"
	"testing"
)

func TestList(t *testing.T) {
	fact, err := NewFactory(house.SourceLeyoujia)
	if err != nil {
		t.Fatal(err)
	}
	param := &crawler.ListParam{
		Page: 1,
	}
	parser := fact.CreateListParser(param)
	list, err := parser.Parse()
	if err != nil {
		t.Fatal(err)
	}
	log.Println(list)
}

// 10350749

func TestSingle(t *testing.T) {
	fact, err := NewFactory(house.SourceLeyoujia)
	if err != nil {
		t.Fatal(err)
	}
	param := &crawler.SingleParam{Id: "9038709"}
	parser := fact.CreateSingleParser(param)
	single, err := parser.Parse()
	if err != nil {
		t.Fatal(err)
	}
	log.Println(single)
}

func TestAll(t *testing.T) {
	log.Println("开始===")
	fact, err := NewFactory(house.SourceLeyoujia)
	if err != nil {
		t.Fatal(err)
	}
	listParam := crawler.NewListParam()
	listParam.Page = 4
	listParser := fact.CreateListParser(listParam)

	list, err := listParser.Parse()
	if err != nil {
		t.Fatal(err)
	}

	count := 0
	total := len(list)
	houses := make([]house.House, 0, len(list))
	for _, item := range list {
		count++
		log.Printf("single parsing %d/%d\n", count, total)
		singleParam := crawler.NewSingleParam()
		singleParam.Id = item.Id
		singleParser := fact.CreateSingleParser(singleParam)
		h, err := singleParser.Parse()
		if err != nil {
			log.Println("single parser: ", err)
		}
		houses = append(houses, *h)
	}

	log.Println(houses)
}
