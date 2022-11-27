package main

import (
	"bigCitySmallHouse/mongodb"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	opts := mongodb.NewOptions()
	opts.Uri = "mongodb://admin:admin@43.138.174.42:27017/"
	err := mongodb.NewDB().ConnectMongodb(opts)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	engine := gin.Default()
	apiGroup := engine.Group("/api")
	{
		apiGroup.POST("/upload_images", UploadImages)
		apiGroup.GET("/image/:filename", ReadImage)
	}
	err := engine.Run(":10003")
	if err != nil {
		log.Fatalln(err)
	}
}
