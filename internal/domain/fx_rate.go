package domain

type FxRate struct {
	currency  string
	symbol    string
	rateToEUR float64
}

func NewFxRate(currency, symbol string, rateToEUR float64) *FxRate {
	return &FxRate{
		currency:  currency,
		symbol:    symbol,
		rateToEUR: rateToEUR,
	}
}

func (fx *FxRate) Currency() string {
	return fx.currency
}

func (fx *FxRate) Symbol() string {
	return fx.symbol
}

func (fx *FxRate) RateToEUR() float64 {
	return fx.rateToEUR
}
