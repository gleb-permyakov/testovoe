package handler

import (
	"io"
	"main/internal/client"
	"main/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Route(client *client.Client) *gin.Engine {
	gin.DefaultWriter = io.Discard
	r := gin.Default()

	r.POST("/", middleware.GiveClient(client), BaseInput)

	return r
}
