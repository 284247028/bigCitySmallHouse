package crawler_base

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

type CrawlerBase struct {
	Req        *http.Request
	HttpClient *http.Client
	Doc        *goquery.Document
}

func GetCrawlerBase() *CrawlerBase {
	return &CrawlerBase{}
}

func (receiver *CrawlerBase) SetReq(req *http.Request) {
	receiver.Req = req
}

func (receiver *CrawlerBase) SetHttpClient(httpClient *http.Client) {
	receiver.HttpClient = httpClient
}

func (receiver *CrawlerBase) SetHeader() {
}

func (receiver CrawlerBase) SetDoc() error {
	var err error
	resp, err := receiver.HttpClient.Do(receiver.Req)
	if err != nil {
		return err
	}
	receiver.Doc, err = goquery.NewDocumentFromReader(resp.Body)
	return nil
}
