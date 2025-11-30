package client

import (
	"net/http"
	"time"
)

type Client struct {
	http *http.Client
}

func New() *Client {
	return &Client{
		http: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c Client) Get(url string) string {
	resp, err := c.http.Get(url)
	if err != nil {
		return "not available"
	}
	defer resp.Body.Close()

	return "available"
}
