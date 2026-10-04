package postgres

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID              uuid.UUID
	PortfolioID     uuid.UUID
	AssetID         uuid.UUID
	TransactionType string
	Quantity        float64
	UnitPrice       float64
	Currency        string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func NewTransaction(
	id uuid.UUID,
	portfolioID uuid.UUID,
	assetID uuid.UUID,
	transactionType string,
	quantity float64,
	unitPrice float64,
	currency string,
	createdAt time.Time,
	updatedAt time.Time,
) *Transaction {
	return &Transaction{
		ID:              id,
		PortfolioID:     portfolioID,
		AssetID:         assetID,
		TransactionType: transactionType,
		Quantity:        quantity,
		UnitPrice:       unitPrice,
		Currency:        currency,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
	}
}
