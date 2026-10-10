package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	BuyTransaction  = "BUY"
	SellTransaction = "SELL"
)

type Position struct {
	id          uuid.UUID
	portfolioID uuid.UUID
	assetID     uuid.UUID
	quantity    float64
	currency    string
	createdAt   time.Time
	updatedAt   time.Time
}

func NewPosition(
	id uuid.UUID,
	portfolioID uuid.UUID,
	assetID uuid.UUID,
	quantity float64,
	currency string,
	createdAt,
	updatedAt time.Time,
) *Position {
	return &Position{
		id:          id,
		portfolioID: portfolioID,
		assetID:     assetID,
		quantity:    quantity,
		currency:    currency,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
	}
}

func (p *Position) ID() uuid.UUID {
	return p.id
}

func (p *Position) PortfolioID() uuid.UUID {
	return p.portfolioID
}

func (p *Position) AssetID() uuid.UUID {
	return p.assetID
}

func (p *Position) Quantity() float64 {
	return p.quantity
}

func (p *Position) Currency() string {
	return p.currency
}

func (p *Position) CreatedAt() time.Time {
	return p.createdAt
}

func (p *Position) UpdatedAt() time.Time {
	return p.updatedAt
}

func (p *Position) Apply(transactionType string, quantity float64, updateTime time.Time) error {
	switch transactionType {
	case BuyTransaction:
		p.quantity += quantity
	case SellTransaction:
		if p.quantity < quantity {
			return errors.New("unsufficient ammount of assets")
		}
		p.quantity -= quantity
	default:
		return errors.New("transaction type not supported")
	}
	p.updatedAt = updateTime
	return nil
}
