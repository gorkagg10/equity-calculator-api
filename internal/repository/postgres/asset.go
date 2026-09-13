package postgres

import (
	"time"

	"github.com/google/uuid"
)

type Asset struct {
	ID           uuid.UUID
	Name         string
	Symbol       string
	Currency     string
	ExchangeName string
	Price        float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewAsset(
	id uuid.UUID,
	name,
	symbol,
	currency,
	exchangeName string,
	price float64,
	createdAt,
	updatedAt time.Time,
) *Asset {
	return &Asset{
		ID:           id,
		Name:         name,
		Symbol:       symbol,
		Currency:     currency,
		ExchangeName: exchangeName,
		Price:        price,
		CreatedAt:    createdAt,
		UpdatedAt:    updatedAt,
	}
}
