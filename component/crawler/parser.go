package crawler

import (
	"io"
	"io/ioutil"
	"net/http"
)

type Parser struct {
	//Req        *http.Request
	//HttpClient *http.Client
}

func NewParser() *Parser {
	return &Parser{
		//Req:        &http.Request{},
		//HttpClient: http.DefaultClient,
	}
}

func (receiver *Parser) Do(client *http.Client, req *http.Request) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	return ioutil.ReadAll(resp.Body)
}
