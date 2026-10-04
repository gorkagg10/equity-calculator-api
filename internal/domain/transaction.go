package domain

import "github.com/google/uuid"

type Transaction struct {
	id              uuid.UUID
	portfolioID     uuid.UUID
	assetID         uuid.UUID
	transactionType string
	quantity        float64
	unitPrice       float64
	currency        string
}

func NewTransaction(
	id uuid.UUID,
	portfolioID uuid.UUID,
	assetID uuid.UUID,
	transactionType string,
	quantity float64,
	unitPrice float64,
	currency string,
) *Transaction {
	return &Transaction{
		id:              id,
		portfolioID:     portfolioID,
		assetID:         assetID,
		transactionType: transactionType,
		quantity:        quantity,
		unitPrice:       unitPrice,
		currency:        currency,
	}
}

func (t *Transaction) ID() uuid.UUID {
	return t.id
}

func (t *Transaction) PortfolioID() uuid.UUID {
	return t.portfolioID
}

func (t *Transaction) AssetID() uuid.UUID {
	return t.assetID
}

func (t *Transaction) Type() string {
	return t.transactionType
}

func (t *Transaction) Quantity() float64 {
	return t.quantity
}

func (t *Transaction) UnitPrice() float64 {
	return t.unitPrice
}

func (t *Transaction) Currency() string {
	return t.currency
}
