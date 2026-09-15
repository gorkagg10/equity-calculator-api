package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/gorkagg10/equity-calculator-api/internal/domain"
)

type Transaction struct {
	transactionRepository domain.TransactionRepository
}

func NewTransaction(
	transactionRepository domain.TransactionRepository,
) *Transaction {
	return &Transaction{
		transactionRepository: transactionRepository,
	}
}

func (t *Transaction) Add(ctx context.Context, symbol string, shares float64) (*domain.Transaction, error) {
	transaction := domain.NewTransaction(uuid.New(), symbol, shares)
	if err := t.transactionRepository.Add(ctx, transaction); err != nil {
		return nil, err
	}

	return transaction, nil
}
