package news

import (
	"net/http"
)

type Client struct {
	http     *http.Client
	key      string
	PageSize int
}

// Create and return a new Client instance for making requests to the News api
// and return max 100 pages
func NewClient(httpClient *http.Client, key string, pageSize int) *Client {
	if pageSize > 100 {
		pageSize = 100
	}

	return &Client{httpClient, key, pageSize}
}
