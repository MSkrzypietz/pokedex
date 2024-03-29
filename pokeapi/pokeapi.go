package pokeapi

import (
	"errors"
	"fmt"
	"github.com/MSkrzypietz/pokedex/pokecache"
	"io"
	"net/http"
	"time"
)

const baseApiUrl = "https://pokeapi.co/api/v2/"

type Client struct {
	httpClient http.Client
	cache      *pokecache.Cache
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}

func (c *Client) httpGet(url string) ([]byte, error) {
	if data, ok := c.cache.Get(url); ok {
		return data, nil
	}

	var data []byte
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return data, errors.New("unable to fetch the location areas")
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return data, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return data, errors.New("couldn't read request")
	}

	c.cache.Add(url, data)
	return data, nil
}

func (c *Client) Close() {
	c.cache.Close()
}
