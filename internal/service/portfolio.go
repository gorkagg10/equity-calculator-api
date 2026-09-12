package service

import (
	"context"

	"github.com/gorkagg10/equity-calculator-api/internal/domain"
)

type Portfolio struct {
	repository domain.PortfolioRepository
}

func NewPortfolio(repository domain.PortfolioRepository) *Portfolio {
	return &Portfolio{
		repository: repository,
	}
}

func (p *Portfolio) AddPortfolio(ctx context.Context, name string) (*domain.Portfolio, error) {
	portfolio, err := domain.NewPortfolio(name)
	if err != nil {
		return nil, err
	}
	if err = p.repository.AddPortfolio(ctx, portfolio); err != nil {
		return nil, err
	}
	return portfolio, nil
}
