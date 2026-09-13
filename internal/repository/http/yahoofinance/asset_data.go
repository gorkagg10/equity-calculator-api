package yahoofinance

type AssetData struct {
	Currency           string  `json:"currency"`
	Symbol             string  `json:"symbol"`
	Name               string  `json:"longName"`
	RegularMarketPrice float64 `json:"regularMarketPrice"`
	RegularMarketTime  int64   `json:"regularMarketTime"`
	ExchangeName       string  `json:"exchangeName"`
}
