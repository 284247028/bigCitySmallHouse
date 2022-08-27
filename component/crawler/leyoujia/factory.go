package leyoujia

import (
	"bigCitySmallHouse/component/crawler"
)

type Factory struct {
}

func NewFactory() crawler.FactoryInterface {
	return &Factory{}
}

func (receiver Factory) CreateSingleParser(param *crawler.SingleParam) crawler.SingleParserInterface {
	return NewSingleParser(param)
}

func (receiver Factory) CreateListParser(param *crawler.ListParam) crawler.ListParserInterface {
	return NewListParser(param)
}
