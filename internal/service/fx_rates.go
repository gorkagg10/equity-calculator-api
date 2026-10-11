package service

import (
	"context"

	"github.com/gorkagg10/equity-calculator-api/internal/domain"
)

type FxRate struct {
	fxRateRepository    domain.FxRateRepository
	assetDataRepository domain.AssetDataRepository
}

func NewFxRate(
	fxRateRepository domain.FxRateRepository,
	assetDataRepository domain.AssetDataRepository,
) *FxRate {
	return &FxRate{
		fxRateRepository:    fxRateRepository,
		assetDataRepository: assetDataRepository,
	}
}

func (fx *FxRate) RefreshRate(ctx context.Context, currency, symbol string) error {
	assetData, err := fx.assetDataRepository.GetAssetData(symbol)
	if err != nil {
		return err
	}

	fxRate := domain.NewFxRate(currency, symbol, assetData.Price())

	return fx.fxRateRepository.Upsert(ctx, fxRate)
}
