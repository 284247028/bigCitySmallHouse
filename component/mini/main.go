package main

import (
	"bigCitySmallHouse/component/crawler/model/house"
	"bigCitySmallHouse/mongodb"
	"bigCitySmallHouse/mongodb/collection"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

func init() {
	err := mongodb.NewDB().ConnectMongodb(mongodb.NewOptions())
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	engine := gin.Default()

	engine.GET("/house_all", allHouse)

	err := engine.Run(":10000")
	if err != nil {
		log.Fatalln(err)
	}
}

func allHouse(ctx *gin.Context) {
	opts := collection.NewOptions()
	opts.DB = mongodb.DBCrawler
	opts.Collection = house.CollectionName
	coll := collection.NewCollection(opts)

	var houses []house.House
	findOpts := &options.FindOptions{}
	findOpts.SetLimit(50)

	cursor, err := coll.MCollection().Find(context.TODO(), bson.D{}, findOpts)
	if err != nil {
		log.Fatalln(err)
	}

	err = cursor.All(context.TODO(), &houses)
	if err != nil {
		log.Fatalln(err)
	}

	ctx.JSON(http.StatusOK, houses)
}
