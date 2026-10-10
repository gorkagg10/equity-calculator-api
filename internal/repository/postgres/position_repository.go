package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"github.com/gorkagg10/equity-calculator-api/internal/domain"
)

type PositionRepository struct {
	pgClient *sql.DB
}

func NewPositionRepository(pgClient *sql.DB) *PositionRepository {
	return &PositionRepository{
		pgClient: pgClient,
	}
}

func (p PositionRepository) Upsert(ctx context.Context, position *domain.Position) error {
	if err := p.pgClient.QueryRowContext(
		ctx,
		`INSERT INTO positions (id, portfolio_id, asset_id, quantity, currency, created_at, updated_at)
			VALUES($1, $2, $3, $4, $5, $6, $7)
		 ON CONFLICT (portfolio_id, asset_id)
		 DO UPDATE SET
		 	quantity = EXCLUDED.quantity,
			updated_at = EXCLUDED.updated_at;
		`,
		position.ID(), position.PortfolioID(), position.AssetID(), position.Quantity(), position.Currency(), position.CreatedAt(), position.UpdatedAt(),
	).Err(); err != nil {
		return err
	}
	return nil
}

func (p PositionRepository) LoadPosition(ctx context.Context, portfolioID, assetID uuid.UUID) (*domain.Position, error) {
	var position Position

	if err := p.pgClient.QueryRowContext(
		ctx,
		`SELECT id, portfolio_id, asset_id, quantity, currency, created_at, updated_at
		 FROM positions
		 WHERE portfolio_id = $1 AND asset_id = $2
		`, portfolioID, assetID).
		Scan(&position.ID, &position.PortfolioID, &position.AssetID, &position.Quantity, &position.Currency, &position.CreatedAt, &position.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrPositionNotFound
		}
		return nil, err
	}

	return domain.NewPosition(
		position.ID,
		position.PortfolioID,
		position.AssetID,
		position.Quantity,
		position.Currency,
		position.CreatedAt,
		position.UpdatedAt,
	), nil
}
