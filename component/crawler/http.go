package crawler

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

var ReqNilErr = errors.New("req nil")
var HttpClientErr = errors.New("httpClient nil")

type HttpCrawler struct {
	*Crawler
	Req        *http.Request
	HttpClient *http.Client
}

func NewHttpCrawler() *HttpCrawler {
	crawler := NewCrawler()
	return &HttpCrawler{Crawler: crawler}
}

func (receiver *HttpCrawler) Fetch() ([]byte, error) {
	if receiver.Req == nil {
		return nil, ReqNilErr
	}
	if receiver.HttpClient == nil {
		return nil, HttpClientErr
	}

	resp, err := receiver.HttpClient.Do(receiver.Req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	return ioutil.ReadAll(resp.Body)
}
