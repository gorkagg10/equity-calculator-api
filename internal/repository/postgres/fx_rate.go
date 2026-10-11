package postgres

import "time"

type FxRate struct {
	Currency  string
	Symbol    string
	RateToEUR float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewFxRate(
	currency,
	symbol string,
	rateToEUR float64,
	createdAt,
	updatedAt time.Time,
) *FxRate {
	return &FxRate{
		Currency:  currency,
		Symbol:    symbol,
		RateToEUR: rateToEUR,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
