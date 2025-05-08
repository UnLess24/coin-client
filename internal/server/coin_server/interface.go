package coinserver

import (
	"context"
	"fmt"
)

type CoinServer interface {
	Close() error
	List(ctx context.Context) ([]Currency, error)
	Currency(ctx context.Context, currency string) (*QuoteResponse, error)
}

func New(schema, host, port, srvType string) (CoinServer, error) {
	switch srvType {
	case "http":
		return NewHTTPServer(schema, host, port), nil
	case "grpc":
		return NewGRPCServer(host, port)
	default:
		return nil, fmt.Errorf("unknown server type: %s", srvType)
	}
}
