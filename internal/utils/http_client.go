package utils

import (
	"context"
	"io/ioutil"
	"net/http"
)

type HTTPClientInterface interface {
	Get(ctx context.Context, url string) ([]byte, error)
}

type HTTPClient struct {
	httpClient *http.Client
}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		httpClient: &http.Client{},
	}
}

func (c *HTTPClient) Get(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	return data, nil
}
