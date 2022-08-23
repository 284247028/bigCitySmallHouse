package crawler

import "bigCitySmallHouse/model"

type Parser interface {
	Parse() ([]model.House, error)
}

type Crawler struct {
}

func DefaultCrawler() *Crawler {
	return &Crawler{}
}
