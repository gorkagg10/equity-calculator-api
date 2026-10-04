package domain

import (
	"context"

	"github.com/google/uuid"
)

//go:generate mockgen -destination=mock_repositories.go -package=domain . PortfolioRepository
type PortfolioRepository interface {
	AddPortfolio(ctx context.Context, portfolio *Portfolio) error
	FindByID(ctx context.Context, portfolioID uuid.UUID) (*Portfolio, error)
}

type AssetDataRepository interface {
	GetAssetData(symbol string) (*AssetData, error)
}

type AssetRepository interface {
	Add(ctx context.Context, asset *Asset) error
	FindByID(ctx context.Context, assetID uuid.UUID) (*Asset, error)
}

type TransactionRepository interface {
	Add(ctx context.Context, transaction *Transaction) error
}

type PositionRepository interface {
	LoadPosition(ctx context.Context, portfolioID uuid.UUID, assetID uuid.UUID) (*Position, error)
	Upsert(ctx context.Context, position *Position) error
}
