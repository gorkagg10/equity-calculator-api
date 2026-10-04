package portfolio

type GetPortfolioResponse struct {
	Equities         []Equity `json:"equities"`
	TotalMarketValue float64  `json:"totalMarketValue"`
}

type Equity struct {
	Ticker      string  `json:"ticker"`
	Allocation  float64 `json:"allocation"`
	MarketValue float64 `json:"marketValue"`
	Shares      float64 `json:"shares"`
}
