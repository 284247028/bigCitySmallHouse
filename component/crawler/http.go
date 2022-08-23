package crawler

import (
	"bigCitySmallHouse/model"
)

type HttpCrawler struct {
	*Crawler
}

func DefaultHttpCrawler() *HttpCrawler {
	crawler := DefaultCrawler()
	return &HttpCrawler{Crawler: crawler}
}

func NewHttpCrawler(crawler *Crawler) *HttpCrawler {
	return &HttpCrawler{Crawler: crawler}
}

func (receiver *HttpCrawler) Parse() ([]model.House, error) {
	//TODO implement me
	panic("implement me")
}
