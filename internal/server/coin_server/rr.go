package coinserver

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func (s *CoinServer) response(ctx context.Context, url string) ([]byte, error) {
	client := &http.Client{}
	req, err := s.request(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to send request url: %s, err: %w", url, err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get response url: %s, err: %w", url, err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	return body, nil
}

func (s *CoinServer) request(ctx context.Context, url string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s%s", s.addr, url), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accepts", "application/json")

	return req, nil
}
