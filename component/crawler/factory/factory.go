package factory

import (
	"bigCitySmallHouse/component/crawler"
	"bigCitySmallHouse/component/crawler/anjuke"
	"bigCitySmallHouse/component/crawler/leyoujia"
	"bigCitySmallHouse/component/crawler/model/house"
	"fmt"
)

func NewFactory(source house.Source) (crawler.FactoryInterface, error) {
	switch source {
	case house.SourceLeyoujia:
		return leyoujia.NewFactory(), nil
	case house.SourceAnjuke:
		return anjuke.NewFactory(), nil
	default:
		return nil, fmt.Errorf("source error: %s", source)
	}
}
