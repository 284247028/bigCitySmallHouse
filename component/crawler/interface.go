package crawler

import (
	"bigCitySmallHouse/model/house"
)

type FactoryInterface interface {
	CreateSingleParser(param *SingleParam) SingleParserInterface
	CreateListParser(param *ListParam) ListParserInterface
}

type ParserInterface interface {
}

type ListParserInterface interface {
	ParserInterface
	Parse() ([]house.House, error)
	Info() (*ListInfo, error)
}

type SingleParserInterface interface {
	ParserInterface
	Parse() (*house.House, error)
}
