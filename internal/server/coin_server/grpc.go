package coinserver

import (
	"context"
	"fmt"

	"github.com/UnLess24/coin/client/pkg/api/coin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ CoinServer = (*grpcCoinServer)(nil)

func NewGRPCServer(host, port string) (*grpcCoinServer, error) {
	url := fmt.Sprintf("%s:%s", host, port)
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to grpc server: %w", err)
	}

	client := coin.NewCoinServiceClient(conn)
	return &grpcCoinServer{client: client, conn: conn}, nil
}

type grpcCoinServer struct {
	client coin.CoinServiceClient
	conn   *grpc.ClientConn
}

func (s *grpcCoinServer) Close() error {
	return s.conn.Close()
}

func (s *grpcCoinServer) List(ctx context.Context) ([]Currency, error) {
	resp, err := s.client.ListCoins(ctx, &coin.ListCoinsRequest{})
	if err != nil {
		return nil, err
	}

	var currencies []Currency
	for _, currency := range resp.GetCurrencies() {
		quote := make(map[string]Quote)
		for k, v := range currency.GetQuote() {
			quote[k] = Quote{
				Price:       v.GetPrice(),
				LastUpdated: v.GetLastUpdated(),
			}
		}
		currencies = append(currencies, Currency{
			ID:     currency.GetId(),
			Name:   currency.GetName(),
			Symbol: currency.GetSymbol(),
			Quote:  quote,
		})
	}
	return currencies, nil
}

func (s *grpcCoinServer) Currency(ctx context.Context, currency string) (*QuoteResponse, error) {
	resp, err := s.client.GetQuote(ctx, &coin.GetQuoteRequest{Name: currency})
	if err != nil {
		return nil, err
	}

	var quotes []Quote
	for _, quote := range resp.GetQuotes() {
		quotes = append(quotes, Quote{
			Price:       quote.GetPrice(),
			LastUpdated: quote.GetLastUpdated(),
		})
	}
	return &QuoteResponse{Name: currency, Quotes: quotes}, nil
}
