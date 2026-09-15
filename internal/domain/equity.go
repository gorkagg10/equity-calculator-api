package domain

import "github.com/google/uuid"

type Equity struct {
	id          uuid.UUID
	portfolioID uuid.UUID
	assetID     uuid.UUID
	shares      float64
}

func NewEquity(
	id uuid.UUID,
	portfolioID uuid.UUID,
	assetID uuid.UUID,
	shares float64,
) {

}
