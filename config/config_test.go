package config

import (
	"log"
	"testing"
)

func TestPath(t *testing.T) {
	path, err := CurrencyPath()
	if err != nil {
		t.Fatal(err)
	}
	log.Println(path)
}
