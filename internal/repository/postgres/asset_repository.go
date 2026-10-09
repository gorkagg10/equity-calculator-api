package postgres

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/google/uuid"

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

func (a *AssetRepository) Add(ctx context.Context, domainAsset *domain.Asset) error {
	asset := NewAsset(
		domainAsset.ID(),
		domainAsset.Name(),
		domainAsset.Symbol(),
		domainAsset.Currency(),
		domainAsset.Exchange(),
		domainAsset.Price(),
		domainAsset.CreatedAt(),
		domainAsset.UpdatedAt(),
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

func (a *AssetRepository) FindByID(ctx context.Context, assetID uuid.UUID) (*domain.Asset, error) {
	var asset Asset

	if err := a.pgClient.QueryRowContext(
		ctx,
		`SELECT id, name, symbol, currency, exchange, price, created_at, updated_at
		 FROM assets
		 WHERE id = $1
		`, assetID).Scan(&asset.ID, &asset.Name, &asset.Symbol, &asset.Currency, &asset.ExchangeName, &asset.Price, &asset.CreatedAt, &asset.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrAssetNotFound
		}
		return nil, err
	}

	domainAsset, err := domain.NewAsset(
		asset.ID,
		domain.NewAssetData(
			asset.Symbol,
			asset.Name,
			asset.Price,
			asset.Currency,
			asset.ExchangeName,
		),
		asset.CreatedAt,
		asset.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return domainAsset, nil
}
