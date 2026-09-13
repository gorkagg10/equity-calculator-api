package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/gorkagg10/equity-calculator-api/internal/domain"
)

type Asset struct {
	assetRepository     domain.AssetRepository
	assetDataRepository domain.AssetDataRepository
}

func NewAsset(
	assetRepository domain.AssetRepository,
	assetDataRepository domain.AssetDataRepository) *Asset {
	return &Asset{
		assetRepository:     assetRepository,
		assetDataRepository: assetDataRepository,
	}
}

func (a *Asset) Add(ctx context.Context, symbol string) (*domain.Asset, error) {
	assetData, err := a.assetDataRepository.GetAssetData(symbol)
	if err != nil {
		return nil, err
	}
	asset, err := domain.NewAsset(
		uuid.New(),
		assetData,
	)
	if err = a.assetRepository.Add(ctx, asset); err != nil {
		return nil, err
	}
	return asset, nil
}
