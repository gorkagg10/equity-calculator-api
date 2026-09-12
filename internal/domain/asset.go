package domain

import (
	"time"

	"github.com/google/uuid"
)

type Asset struct {
	id uuid.UUID
	AssetData
	createdAt time.Time
	updatedAt time.Time
}

func NewAsset(
	id uuid.UUID,
	assetData AssetData,
	createdAt,
	updatedAt time.Time,
) (*Asset, error) {
	return &Asset{
		id:        id,
		AssetData: assetData,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}
