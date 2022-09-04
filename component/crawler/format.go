package crawler

import (
	"bigCitySmallHouse/model/house"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func FormatHouses(houses []house.House) []interface{} {
	fHouses := make([]interface{}, 0, len(houses))
	for _, h := range houses {
		fHouse, err := bson.Marshal(h)
		if err != nil {
			log.Println(err)
			continue
		}
		fHouses = append(fHouses, fHouse)
	}
	return fHouses
}
