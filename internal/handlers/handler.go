package handlers

import (
	"io"

	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	gin.DefaultWriter = io.Discard
	r := gin.Default()

	r.POST("/", BaseInput)

	return r
}
