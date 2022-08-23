package crawler

import (
	"bigCitySmallHouse/model"
	"github.com/PuerkitoBio/goquery"
)

type HtmlCrawler struct {
	*Crawler
	Doc *goquery.Document
}

func (receiver *HtmlCrawler) Parse() ([]model.House, error) {
	//TODO implement me
	panic("implement me")
}
