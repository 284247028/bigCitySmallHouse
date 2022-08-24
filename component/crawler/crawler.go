package crawler

import "bigCitySmallHouse/model"

type Parser interface {
	Parse() ([]model.House, error)
}

type Crawler struct {
	Page int
}

func NewCrawler() *Crawler {
	return &Crawler{}
}
