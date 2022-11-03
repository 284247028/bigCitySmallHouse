package main

import "github.com/gin-gonic/gin"

type BaseResponse struct {
	ctx *gin.Context
}

func NewBaseResponse(ctx *gin.Context) *BaseResponse {
	return &BaseResponse{ctx: ctx}
}

type ErrorMessage struct {
	HttpCode     int    `json:"http_code"`
	ErrorMessage string `json:"error_message"`
}

func (receiver *BaseResponse) ErrorResponse(httpCode int, err error) {
	receiver.ctx.JSON(httpCode, ErrorMessage{
		HttpCode:     httpCode,
		ErrorMessage: err.Error(),
	})
}
