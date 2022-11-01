package main

import "github.com/gin-gonic/gin"

type Router struct {
	*gin.Engine
}

func NewRouter(engine *gin.Engine) *Router {
	return &Router{Engine: engine}
}

func (receiver *Router) Group() {

}
