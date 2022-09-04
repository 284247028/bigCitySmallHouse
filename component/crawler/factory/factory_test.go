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
	info, err := parser.Info()
	if err != nil {
		t.Fatal(err)
	}
	list, err := parser.Parse()
	if err != nil {
		t.Fatal(err)
	}
	log.Println(list, info)
}

// 10350749

func TestSingle(t *testing.T) {
	fact, err := NewFactory(house.SourceLeyoujia)
	if err != nil {
		t.Fatal(err)
	}
	param := &crawler.SingleParam{Id: "90920524"}
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
		singleParam.Id = item.SourceId
		singleParser := fact.CreateSingleParser(singleParam)
		h, err := singleParser.Parse()
		if err != nil {
			log.Println("single parser: ", err)
		}
		houses = append(houses, *h)
	}

	log.Println(houses)
}

func TestArr(t *testing.T) {
	var arr = []int{1, 2, 3}
	Append(arr)
	//log.Println(arr)
}

func Append(arr []int) {
	for i := 4; i < 20; i++ {
		arr = append(arr, i)
	}
	tMap := make(map[int]int)
	for _, v := range arr {
		tMap[v] = v
	}
	log.Println(tMap)
}
