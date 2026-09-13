package domain

import (
	"github.com/google/uuid"
)

type Asset struct {
	id uuid.UUID
	*AssetData
}

func NewAsset(
	id uuid.UUID,
	assetData *AssetData,
) (*Asset, error) {
	return &Asset{
		id:        id,
		AssetData: assetData,
	}, nil
}

func (a Asset) ID() uuid.UUID {
	return a.id
}
