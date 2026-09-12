package service

import (
	"github.com/gorkagg10/equity-calculator-api/internal/domain"
)

type Portfolio struct {
}

func NewPortfolio() *Portfolio {
	return &Portfolio{}
}

func (p *Portfolio) AddPortfolio(name string) (*domain.Portfolio, error) {
	portfolio, err := domain.NewPortfolio(name)
	if err != nil {
		return nil, err
	}
	return portfolio, nil
}
