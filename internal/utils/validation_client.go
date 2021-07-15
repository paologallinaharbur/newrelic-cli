package utils

import (
	"context"
)

// TODO: RENAME TO ValidationService
// Should ValidationClient live in it's own file (e.g. validation_client.go)
type ValidationClient struct {
	httpClient HTTPClientInterface
}

func NewValidationClient() *ValidationClient {
	return &ValidationClient{
		httpClient: NewHTTPClient(),
	}
}

func (c *ValidationClient) Get(ctx context.Context, url string) ([]byte, error) {
	resp, err := c.httpClient.Get(ctx, url)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
