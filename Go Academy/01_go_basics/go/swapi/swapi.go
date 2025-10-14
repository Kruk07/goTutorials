package swapi

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	http *resty.Client
}

func New() *Client {
	return &Client{
		http: resty.New().SetBaseURL("https://swapi.dev/api"),
	}
}

func (c *Client) CharacterExists(name string) (bool, error) {
	resp, err := c.http.R().
		SetQueryParam("search", name).
		Get("/people/")
	if err != nil {
		return false, err
	}

	var result struct {
		Count int `json:"count"`
	}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return false, err
	}

	return result.Count > 0, nil
}
