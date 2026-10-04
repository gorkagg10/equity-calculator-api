package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidName = errors.New("name must not be empty")
)

type Portfolio struct {
	id   uuid.UUID
	name string
}

func NewPortfolio(
	id uuid.UUID,
	name string,
) (*Portfolio, error) {
	if name == "" {
		return nil, ErrInvalidName
	}
	return &Portfolio{
		id:   id,
		name: name,
	}, nil
}

func (p Portfolio) ID() uuid.UUID {
	return p.id
}

func (p Portfolio) Name() string {
	return p.name
}
