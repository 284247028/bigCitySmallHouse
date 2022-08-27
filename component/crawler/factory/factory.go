package factory

import (
	"bigCitySmallHouse/component/crawler"
	"bigCitySmallHouse/component/crawler/leyoujia"
	"bigCitySmallHouse/model/house"
	"errors"
)

func NewFactory(source house.Source) (crawler.FactoryInterface, error) {
	switch source {
	case house.SourceLeyoujia:
		return leyoujia.NewFactory(), nil
	default:
		return nil, errors.New("new crawler factory")
	}
}
