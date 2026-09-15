package postgres

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

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
		domainTransaction.Symbol(),
		domainTransaction.Shares(),
		time.Now(),
		time.Now(),
	)

	if err := t.pgClient.QueryRowContext(
		ctx,
		`INSERT INTO transactions (id, symbol, shares, created_at, updated_at)
				VALUES($1, $2, $3, $4, $5);`,
		transaction.ID,
		transaction.Symbol,
		transaction.Shares,
		transaction.CreatedAt,
		transaction.UpdatedAt,
	).Err(); err != nil {
		slog.ErrorContext(ctx, "inserting transaction in database", slog.String("error", err.Error()))
		return err
	}
	return nil
}
