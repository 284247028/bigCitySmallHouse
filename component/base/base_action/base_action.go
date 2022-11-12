package base_action

import "github.com/gin-gonic/gin"

type ErrorMessage struct {
	HttpCode     int    `json:"http_code"`
	ErrorMessage string `json:"error_message"`
}

func ErrorResponse(ctx *gin.Context, httpCode int, err error) {
	ctx.JSON(httpCode, ErrorMessage{
		HttpCode:     httpCode,
		ErrorMessage: err.Error(),
	})
}
