package domain

import (
	"time"

	"github.com/google/uuid"
)

type Asset struct {
	id uuid.UUID
	*AssetData
	createdAt time.Time
	updatedAt time.Time
}

func NewAsset(
	id uuid.UUID,
	assetData *AssetData,
	createdAt time.Time,
	updatedAt time.Time,
) (*Asset, error) {
	return &Asset{
		id:        id,
		AssetData: assetData,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

func (a Asset) ID() uuid.UUID {
	return a.id
}

func (a Asset) CreatedAt() time.Time {
	return a.createdAt
}

func (a Asset) UpdatedAt() time.Time {
	return a.updatedAt
}
