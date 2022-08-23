package leyoujia

import (
	"bigCitySmallHouse/component/crawler"
	"bigCitySmallHouse/model"
	"crypto/md5"
	"fmt"
)

const Domain = "https://steward.leyoujia.com"
const Path = "/stewardnew/zf/queryZfList"

type Crawler struct {
	*crawler.HttpCrawler
}

func DefaultCrawler() *Crawler {
	httpCrawler := crawler.DefaultHttpCrawler()
	return NewCrawler(httpCrawler)
}

func NewCrawler(httpCrawler *crawler.HttpCrawler) *Crawler {
	return &Crawler{HttpCrawler: httpCrawler}
}

func (receiver *Crawler) Parse() ([]model.House, error) {
	//TODO implement me
	panic("implement me")
}

func sign(str string) string {
	res := md5.Sum([]byte(str))
	val := fmt.Sprintf("%x", res)
	bytes := []byte(val) // 这是32字节
	res = md5.Sum(bytes)
	val = fmt.Sprintf("%x", res)
	return val
}
