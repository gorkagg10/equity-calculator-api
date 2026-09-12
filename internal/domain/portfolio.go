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

func NewPortfolio(name string) (*Portfolio, error) {
	if name == "" {
		return nil, ErrInvalidName
	}
	return &Portfolio{
		id:        uuid.New(),
		name:      name,
		createdAt: time.Now().UTC(),
		updatedAt: time.Now().UTC(),
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
