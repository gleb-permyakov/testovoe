package handlers

import (
	"log"
	"main/internal/delivery/http/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BaseInput(c *gin.Context) {

	var body struct {
		Links []string `json:"links" binding:"required"`
	}

	err := c.Bind(&body)
	if err != nil {
		code := 400
		msg := "invalid body"
		response.ResponseError(c, code, msg)
		log.Printf("error on POST /: %d %s", code, msg)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": body.Links,
	})
}
