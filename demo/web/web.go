package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/", index)
	err := engine.Run(":10000")
	if err != nil {
		log.Fatalln(err)
	}

}

func index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, struct {
		Msg string `json:"msg"`
	}{
		"hello world!",
	})
}
