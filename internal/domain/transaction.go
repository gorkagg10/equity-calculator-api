package domain

import "github.com/google/uuid"

type Transaction struct {
	id     uuid.UUID
	symbol string
	shares float64
}

func (t *Transaction) ID() uuid.UUID {
	return t.id
}

func (t *Transaction) Symbol() string {
	return t.symbol
}

func (t *Transaction) Shares() float64 {
	return t.shares
}

func NewTransaction(
	id uuid.UUID,
	symbol string,
	shares float64,
) *Transaction {
	return &Transaction{
		id:     id,
		symbol: symbol,
		shares: shares,
	}
}
