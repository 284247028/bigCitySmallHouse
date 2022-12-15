package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	engine.Any("/:serviceName/*proxyPath", apiProxy)
	err := engine.Run(":11000")
	if err != nil {
		log.Fatalln(err)
	}
}
