package postgres

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/gorkagg10/equity-calculator-api/internal/domain"
)

type TransactionRepository struct {
	pgClient *sql.DB
}

func NewTransactionRepository(pgClient *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		pgClient: pgClient,
	}
}

func (t TransactionRepository) Add(ctx context.Context, domainTransaction *domain.Transaction) error {
	transaction := NewTransaction(
		domainTransaction.ID(),
		domainTransaction.PortfolioID(),
		domainTransaction.AssetID(),
		domainTransaction.Type(),
		domainTransaction.Quantity(),
		domainTransaction.UnitPrice(),
		domainTransaction.Currency(),
		domainTransaction.CreatedAt(),
		domainTransaction.UpdatedAt(),
	)

	if err := t.pgClient.QueryRowContext(
		ctx,
		`INSERT INTO transactions (id, portfolio_id, asset_id, transaction_type, quantity, unit_price, currency, created_at, updated_at)
				VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);`,
		transaction.ID,
		transaction.PortfolioID,
		transaction.AssetID,
		transaction.TransactionType,
		transaction.Quantity,
		transaction.UnitPrice,
		transaction.Currency,
		transaction.CreatedAt,
		transaction.UpdatedAt,
	).Err(); err != nil {
		slog.ErrorContext(ctx, "inserting transaction in database", slog.String("error", err.Error()))
		return err
	}
	return nil
}
