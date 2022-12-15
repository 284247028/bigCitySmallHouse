package base_action

import "github.com/gin-gonic/gin"

type ErrorMessage struct {
	//HttpCode     int    `json:"http_code"`
	Message string `json:"message"`
}

func ErrorResponse(ctx *gin.Context, httpCode int, err error) {
	ctx.JSON(httpCode, ErrorMessage{
		//HttpCode:     httpCode,
		Message: err.Error(),
	})
}
