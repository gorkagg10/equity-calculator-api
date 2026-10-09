package service

import (
	"context"
	"time"

	"github.com/google/uuid"
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
	portfolio, err := domain.NewPortfolio(
		uuid.New(),
		name,
		time.Now().UTC(),
		time.Now().UTC(),
	)
	if err != nil {
		return nil, err
	}
	if err = p.repository.AddPortfolio(ctx, portfolio); err != nil {
		return nil, err
	}
	return portfolio, nil
}
