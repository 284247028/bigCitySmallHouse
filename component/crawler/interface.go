package crawler

import (
	"bigCitySmallHouse/component/crawler/model/house"
)

type FactoryInterface interface {
	CreateSingleParser(param *SingleParam) SingleParserInterface
	CreateListParser(param *ListParam) ListParserInterface
}

type ParserInterface interface {
}

type ListParserInterface interface {
	ParserInterface
	Parse() ([]house.House, *ListInfo, error)
}

type SingleParserInterface interface {
	ParserInterface
	Parse() (*house.House, error)
}
