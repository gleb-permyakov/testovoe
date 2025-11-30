package middleware

import (
	"main/internal/client"

	"github.com/gin-gonic/gin"
)

func GiveClient(client *client.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("client", client)
		c.Next()
	}
}
