package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/gorkagg10/equity-calculator-api/internal/domain"
)

const (
	EmptyPositionQuantity = 0
)

type Transaction struct {
	transactionRepository domain.TransactionRepository
	portfolioRepository   domain.PortfolioRepository
	assetRepository       domain.AssetRepository
	positionRepository    domain.PositionRepository
}

func NewTransaction(
	transactionRepository domain.TransactionRepository,
	portfolioRepository domain.PortfolioRepository,
	assetRepository domain.AssetRepository,
	positionRepository domain.PositionRepository,
) *Transaction {
	return &Transaction{
		transactionRepository: transactionRepository,
		portfolioRepository:   portfolioRepository,
		assetRepository:       assetRepository,
		positionRepository:    positionRepository,
	}
}

func (t *Transaction) Add(
	ctx context.Context,
	portfolioID uuid.UUID,
	assetID uuid.UUID,
	unitPrice,
	quantity float64,
	currency,
	transactionType string,
) (*domain.Transaction, error) {
	_, err := t.portfolioRepository.FindByID(ctx, portfolioID)
	if err != nil {
		return nil, err
	}

	_, err = t.assetRepository.FindByID(ctx, assetID)
	if err != nil {
		return nil, err
	}

	transaction := domain.NewTransaction(
		uuid.New(),
		portfolioID,
		assetID,
		transactionType,
		quantity,
		unitPrice,
		currency,
		time.Now().UTC(),
		time.Now().UTC(),
	)
	if err = t.transactionRepository.Add(ctx, transaction); err != nil {
		return nil, err
	}

	position, err := t.positionRepository.LoadPosition(ctx, portfolioID, assetID)
	if err != nil {
		if errors.Is(err, domain.ErrPositionNotFound) {
			position = domain.NewPosition(
				uuid.New(),
				portfolioID,
				assetID,
				EmptyPositionQuantity,
				transaction.Currency(),
				time.Now().UTC(),
				time.Now().UTC(),
			)
		}
	}

	if err = position.Apply(transaction.Type(), transaction.Quantity(), time.Now().UTC()); err != nil {
		return nil, err
	}

	if err = t.positionRepository.Upsert(ctx, position); err != nil {
		return nil, err
	}

	return transaction, nil
}
