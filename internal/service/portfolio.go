package service

import (
	portfoliodto "github.com/gorkagg10/equity-calculator-api/internal/http/dto/portfolio"
)

type Portfolio struct {
}

func NewPortfolio() *Portfolio {
	return &Portfolio{}
}

func (p *Portfolio) AddPortfolio(addPortfolioRequest *portfoliodto.AddPortfolioRequest) error {
	return nil
}
