package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/gorkagg10/equity-calculator-api/internal/domain"
)

type PortfolioRepository struct {
	pgClient *sql.DB
}

func NewPortfolioRepository(pgClient *sql.DB) *PortfolioRepository {
	return &PortfolioRepository{
		pgClient: pgClient,
	}
}

func (p PortfolioRepository) AddPortfolio(ctx context.Context, portfolio *domain.Portfolio) error {
	if err := p.pgClient.QueryRowContext(
		ctx,
		`INSERT INTO portfolios (id, name, created_at, updated_at)
				VALUES($1, $2, $3, $4);`, portfolio.ID(), portfolio.Name(), time.Now(), time.Now()).Err(); err != nil {
		slog.ErrorContext(ctx, "inserting portfolio in database", slog.String("error", err.Error()))
		return err
	}
	return nil
}

func (p PortfolioRepository) FindByID(ctx context.Context, portfolioID uuid.UUID) (*domain.Portfolio, error) {
	var portfolio Portfolio

	if err := p.pgClient.QueryRowContext(
		ctx,
		`SELECT id, name, created_at, updated_at
		 FROM portfolios
		 WHERE id = $1
		`, portfolioID).Scan(&portfolio.ID, &portfolio.Name, &portfolio.CreatedAt, &portfolio.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("portfolio not found")
		}
		return nil, err
	}

	domainPortfolio, err := domain.NewPortfolio(
		portfolio.ID,
		portfolio.Name,
	)
	if err != nil {
		return nil, err
	}

	return domainPortfolio, nil
}
