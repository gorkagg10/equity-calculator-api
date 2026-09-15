package postgres

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID        uuid.UUID
	Symbol    string
	Shares    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTransaction(
	id uuid.UUID,
	symbol string,
	shares float64,
	createdAt time.Time,
	updatedAt time.Time,
) *Transaction {
	return &Transaction{
		ID:        id,
		Symbol:    symbol,
		Shares:    shares,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
