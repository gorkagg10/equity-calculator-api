package postgres

import (
	"time"

	"github.com/google/uuid"
)

type Position struct {
	ID          uuid.UUID
	PortfolioID uuid.UUID
	AssetID     uuid.UUID
	Quantity    float64
	Currency    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
