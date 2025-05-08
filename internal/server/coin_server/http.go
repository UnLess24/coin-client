package coinserver

import (
	"context"
	"encoding/json"
	"fmt"
)

var _ CoinServer = (*httpCoinServer)(nil)

const (
	listEp     = "/list"
	currencyEp = "/currency"
)

type httpCoinServer struct {
	addr string
}

func NewHTTPServer(schema, host, port string) *httpCoinServer {
	return &httpCoinServer{
		addr: fmt.Sprintf("%s://%s:%s", schema, host, port),
	}
}

func (s *httpCoinServer) Close() error {
	return nil
}

func (s *httpCoinServer) List(ctx context.Context) ([]Currency, error) {
	list, err := s.response(ctx, listEp)
	if err != nil {
		return nil, fmt.Errorf("failed to get response: %w", err)
	}

	response := CurrencyResponse{}
	err = json.Unmarshal(list, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response.Data, nil
}

func (s *httpCoinServer) Currency(ctx context.Context, currency string) (*QuoteResponse, error) {
	resp, err := s.response(ctx, fmt.Sprintf("%s?currency=%s", currencyEp, currency))
	if err != nil {
		return nil, fmt.Errorf("failed to get response: %w", err)
	}

	response := &QuoteResponse{}
	err = json.Unmarshal(resp, response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return response, nil
}
