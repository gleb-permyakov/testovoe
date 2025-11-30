package response

import "github.com/gin-gonic/gin"

func ResponseError(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{
		"error": msg,
	})
}
