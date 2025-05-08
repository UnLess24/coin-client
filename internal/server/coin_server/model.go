package coinserver

type CurrencyResponse struct {
	Data []Currency `json:"data"`
}

type Currency struct {
	ID     int64         `json:"id"`
	Name   string        `json:"name"`
	Symbol string        `json:"symbol"`
	Quote  CurrencyQuote `json:"quote"`
}

type CurrencyQuote map[string]Quote

type Quote struct {
	Price       float32 `json:"price"`
	LastUpdated string  `json:"last_updated"`
}

type QuoteResponse struct {
	Name   string  `json:"name"`
	Quotes []Quote `json:"data"`
}
