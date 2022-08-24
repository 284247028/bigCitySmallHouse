package leyoujia

import (
	"log"
	"testing"
)

func TestFetch(t *testing.T) {
	qCrawler := NewCrawler(1)
	err := qCrawler.Init()
	if err != nil {
		t.Fatal(err)
	}
	bs, err := qCrawler.Fetch()
	if err != nil {
		t.Fatal(err)
	}
	log.Println(bs)
}
