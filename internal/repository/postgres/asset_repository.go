package postgres

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	"github.com/gorkagg10/equity-calculator-api/internal/domain"
)

type AssetRepository struct {
	pgClient *sql.DB
}

func NewAssetRepository(pgClient *sql.DB) *AssetRepository {
	return &AssetRepository{
		pgClient: pgClient,
	}
}

func (a AssetRepository) Add(ctx context.Context, domainAsset *domain.Asset) error {
	asset := NewAsset(
		domainAsset.ID(),
		domainAsset.Name(),
		domainAsset.Symbol(),
		domainAsset.Currency(),
		domainAsset.Exchange(),
		domainAsset.Price(),
		time.Now(),
		time.Now(),
	)

	if err := a.pgClient.QueryRowContext(
		ctx,
		`INSERT INTO assets (id, name, symbol, currency, exchange, price, created_at, updated_at)
				VALUES($1, $2, $3, $4, $5, $6, $7, $8);`,
		asset.ID,
		asset.Name,
		asset.Symbol,
		asset.Currency,
		asset.ExchangeName,
		asset.Price,
		asset.CreatedAt,
		asset.UpdatedAt,
	).Err(); err != nil {
		slog.ErrorContext(ctx, "inserting asset in database", slog.String("error", err.Error()))
		return err
	}
	return nil
}
