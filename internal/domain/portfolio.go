package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidName = errors.New("name must not be empty")
)

type Portfolio struct {
	id        uuid.UUID
	name      string
	createdAt time.Time
	updatedAt time.Time
}

func NewPortfolio(
	id uuid.UUID,
	name string,
	createdAt,
	updatedAt time.Time,
) (*Portfolio, error) {
	if name == "" {
		return nil, ErrInvalidName
	}
	return &Portfolio{
		id:        id,
		name:      name,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}, nil
}

func (p Portfolio) ID() uuid.UUID {
	return p.id
}

func (p Portfolio) Name() string {
	return p.name
}

func (p Portfolio) CreatedAt() time.Time {
	return p.createdAt
}

func (p Portfolio) UpdatedAt() time.Time {
	return p.updatedAt
}
