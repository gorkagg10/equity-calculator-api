package postgres

import (
	"context"
	"database/sql"
	"log/slog"

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
				VALUES($1, $2, $3, $4);`, portfolio.ID(), portfolio.Name(), portfolio.CreatedAt(), portfolio.UpdatedAt()).Err(); err != nil {
		slog.ErrorContext(ctx, "inserting portfolio in database", slog.String("error", err.Error()))
		return err
	}
	return nil
}
