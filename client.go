package fanarttv

import (
	"encoding/json"
	"net/http"
)

// DefaultEndpoint represents the default endpoint of the API
var DefaultEndpoint = "https://webservice.fanart.tv/v3"

// Client represents an API client
type Client struct {
	Endpoint string
	APIKey   string
	Client   *http.Client
}

// New returns a new client
func New(apiKey string) *Client {
	return &Client{
		Endpoint: DefaultEndpoint,
		APIKey:   apiKey,
		Client:   http.DefaultClient,
	}
}

func (c *Client) get(url string, resType interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("api-key", c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	// Decode the JSON response
	if err = json.NewDecoder(resp.Body).Decode(&resType); err != nil {
		return err
	}

	return nil
}
