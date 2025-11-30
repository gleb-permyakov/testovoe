package handler

import (
	"log"
	"main/internal/client"
	"main/internal/delivery/http/response"
	"main/internal/service"
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

	clientIface, exists := c.Get("client")
	if !exists {
		response.ResponseError(c, 400, "error while checkingg links")
		log.Print("error in BaseInput: no client in context")
		return
	}

	clientObj, ok := clientIface.(*client.Client)
	if !ok {
		response.ResponseError(c, 400, "error while checkingg links")
		log.Print("error in BaseInput: failed type assertion")
		return
	}

	service := service.TaskService{Client: clientObj}

	states := service.CheckURL(body.Links)

	c.JSON(http.StatusOK, gin.H{
		"message": states,
	})
}
