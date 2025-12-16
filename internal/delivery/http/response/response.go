package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseError(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"error": msg,
	})
}

func ResponseOK(c *gin.Context, message interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
