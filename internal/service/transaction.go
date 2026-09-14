package service

import (
	"github.com/google/uuid"

	"github.com/gorkagg10/equity-calculator-api/internal/domain"
)

type Transaction struct {
}

func NewTransaction() *Transaction {
	return &Transaction{}
}

func (t *Transaction) Add(symbol string, shares float64) {
	_ = domain.NewTransaction(uuid.New(), symbol, shares)

}
