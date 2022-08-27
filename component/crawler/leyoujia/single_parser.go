package leyoujia

import (
	"bigCitySmallHouse/component/crawler"
	"bigCitySmallHouse/model/house"
)

type SingleParser struct {
	*crawler.SingleParser
}

func NewSingleParser(param *crawler.SingleParam) *SingleParser {
	singleParser := crawler.NewSingleParser(param)
	return &SingleParser{SingleParser: singleParser}
}

func (receiver SingleParser) Parse() (house.House, error) {

}
