package coinserver

import (
	"context"
	"fmt"
)

const (
	listEp     = "/list"
	currencyEp = "/currency/"
)

type CoinServer struct {
	addr string
}

func New(schema, host, port string) *CoinServer {
	return &CoinServer{
		addr: fmt.Sprintf("%s://%s:%s", schema, host, port),
	}
}

func (s *CoinServer) List(ctx context.Context) ([]byte, error) {
	return s.response(ctx, listEp)
}

func (s *CoinServer) Currency(ctx context.Context, currency string) ([]byte, error) {
	return s.response(ctx, fmt.Sprintf("%s%s", currencyEp, currency))
}
