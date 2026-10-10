package service

import (
	"context"
	"time"

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
		time.Now().UTC(),
		time.Now().UTC(),
	)
	if err = a.assetRepository.Add(ctx, asset); err != nil {
		return nil, err
	}
	return asset, nil
}

func (a *Asset) UpdatePrices(ctx context.Context) error {
	assets, err := a.assetRepository.List(ctx)
	if err != nil {
		return err
	}
	for i, asset := range assets {
		assetData, err := a.assetDataRepository.GetAssetData(asset.Symbol())
		if err != nil {
			return err
		}
		assets[i].SetPrice(assetData.Price())

	}
	return nil
}
