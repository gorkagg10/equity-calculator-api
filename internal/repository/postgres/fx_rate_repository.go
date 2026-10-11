package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/gorkagg10/equity-calculator-api/internal/domain"
)

type FxRateRepository struct {
	pgClient *sql.DB
}

func NewFxRateRepository(pgClient *sql.DB) *FxRateRepository {
	return &FxRateRepository{
		pgClient: pgClient,
	}
}

func (fx *FxRateRepository) Upsert(ctx context.Context, domainFxRate *domain.FxRate) error {
	fxRate := NewFxRate(
		domainFxRate.Currency(),
		domainFxRate.Symbol(),
		domainFxRate.RateToEUR(),
		time.Now().UTC(),
		time.Now().UTC(),
	)
	if err := fx.pgClient.QueryRowContext(
		ctx,
		`INSERT INTO fx_rates (currency, symbol, rate_to_eur, created_at, updated_at)
			VALUES($1, $2, $3, $4, $5)
		 ON CONFLICT (currency)
		 DO UPDATE SET
		 	rate_to_eur = EXCLUDED.rate_to_eur,
			updated_at = EXCLUDED.updated_at;
		`,
		fxRate.Currency, fxRate.Symbol, fxRate.RateToEUR, fxRate.CreatedAt, fxRate.UpdatedAt,
	).Err(); err != nil {
		return err
	}
	return nil

}
