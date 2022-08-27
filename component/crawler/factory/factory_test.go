package factory

import (
	"bigCitySmallHouse/component/crawler"
	"bigCitySmallHouse/model/house"
	"log"
	"testing"
)

func TestLeyoujiaList(t *testing.T) {
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
