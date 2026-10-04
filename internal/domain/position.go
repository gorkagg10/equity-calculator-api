package domain

import (
	"errors"

	"github.com/google/uuid"
)

type Position struct {
	id          uuid.UUID
	portfolioID uuid.UUID
	assetID     uuid.UUID
	quantity    float64
	currency    string
}

func NewPosition(
	id uuid.UUID,
	portfolioID uuid.UUID,
	assetID uuid.UUID,
	quantity float64,
	currency string,
) *Position {
	return &Position{
		id:          id,
		portfolioID: portfolioID,
		assetID:     assetID,
		quantity:    quantity,
		currency:    currency,
	}
}

func (p *Position) ApplyTransaction(
	portfolioID,
	assetID uuid.UUID,
	transaction *Transaction,
) error {
	if p == nil {
		p = NewPosition(
			uuid.New(),
			portfolioID,
			assetID,
			transaction.Quantity(),
			transaction.Currency(),
		)
		return nil
	}
	return p.Update(transaction.Type(), transaction.Quantity())
}

func (p *Position) ID() uuid.UUID {
	return p.id
}

func (p *Position) PortfolioID() uuid.UUID {
	return p.portfolioID
}

func (p *Position) AssetID() string {
	return p.AssetID()
}

func (p *Position) Quantity() float64 {
	return p.quantity
}

func (p *Position) Currency() string {
	return p.currency
}

func (p *Position) Update(transactionType string, quantity float64) error {
	switch transactionType {
	case "BUY":
		p.quantity += quantity
	case "SELL":
		if p.quantity < quantity {
			return errors.New("unsufficient ammount of assets")
		}
		p.quantity -= quantity
	default:
		return errors.New("transaction type not supported")
	}
	return nil
}
