package postgres

import (
	"time"

	"github.com/google/uuid"
)

type Portfolio struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
