package pokeapi

import (
	"net/http"
	"pokedex/internal/pokecache"
	"time"
)

type Doer interface { // Enables mocking
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	httpClient Doer
	cache      *pokecache.Cache
	baseURL    string
}

func NewClient(timeout time.Duration, cache *pokecache.Cache) *Client {
	return &Client{
		httpClient: &http.Client{Timeout: timeout},
		cache:      cache,
		baseURL:    baseURL,
	}
}
