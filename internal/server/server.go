package server

import (
	"errors"
	"log"
	"main/internal/config"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	HTTPServer *http.Server
}

func New(cfg *config.Config, r *gin.Engine) *HTTPServer {
	srv := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	return &HTTPServer{
		HTTPServer: srv,
	}
}

func (h HTTPServer) Start(cfg *config.Config) error {
	log.Printf("server started on port %s", cfg.HTTPAddr)
	if err := h.HTTPServer.ListenAndServe(); err != nil {
		return errors.New("err starting server")
	}
	return nil
}
