package domain

import "context"

type PortfolioRepository interface {
	AddPortfolio(ctx context.Context, portfolio *Portfolio) error
}

type AssetDataRepository interface {
	GetAssetData(symbol string) (*AssetData, error)
}

type AssetRepository interface {
	Add(ctx context.Context, asset *Asset) error
}
