package postgres_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/gorkagg10/equity-calculator-api/internal/domain"
	"github.com/gorkagg10/equity-calculator-api/internal/repository/postgres"
)

func TestTransactionRepository_Add(t *testing.T) {
	transactionID := uuid.New()
	portfolioID := uuid.New()
	assetID := uuid.New()
	portfolioTime := time.Now().UTC()
	assetTime := time.Now().UTC()
	now := time.Now().UTC()

	examplePortfolio, err := domain.NewPortfolio(
		portfolioID,
		"test",
		portfolioTime,
		portfolioTime,
	)
	require.NoError(t, err)

	exampleAsset, err := domain.NewAsset(
		assetID,
		domain.NewAssetData(
			"TST",
			"test",
			10.10,
			"EUR",
			"ESP",
		),
		assetTime,
		assetTime,
	)
	require.NoError(t, err)

	exampleTransaction := domain.NewTransaction(
		transactionID,
		portfolioID,
		assetID,
		"BUY",
		10,
		10.10,
		"EUR",
		now,
		now,
	)

	tests := []struct {
		name   string
		before func(
			t *testing.T,
			db *sql.DB,
			transactionRepository *postgres.TransactionRepository,
		)
		domainTransaction *domain.Transaction
		wantErr           bool
	}{
		{
			name:              "portfolio missing error",
			domainTransaction: exampleTransaction,
			wantErr:           true,
		},
		{
			name: "asset missing error",
			before: func(t *testing.T, db *sql.DB, transactionRepository *postgres.TransactionRepository) {
				portfolioRepository := postgres.NewPortfolioRepository(db)
				require.NoError(t, portfolioRepository.AddPortfolio(t.Context(), examplePortfolio))
			},
			domainTransaction: exampleTransaction,
			wantErr:           true,
		},
		{
			name: "transaction created successfully",
			before: func(t *testing.T, db *sql.DB, transactionRepository *postgres.TransactionRepository) {
				portfolioRepository := postgres.NewPortfolioRepository(db)
				require.NoError(t, portfolioRepository.AddPortfolio(t.Context(), examplePortfolio))
				assetRepository := postgres.NewAssetRepository(db)
				require.NoError(t, assetRepository.Add(t.Context(), exampleAsset))
			},
			domainTransaction: exampleTransaction,
		},
	}
	for _, tt := range tests {
		db := startDatabase(t)
		require.NotNil(t, db)
		require.NoError(t, postgres.Migrate(db, "app-test", "migrations"))

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tr := postgres.NewTransactionRepository(db)
			if tt.before != nil {
				tt.before(t, db, tr)
			}
			err := tr.Add(t.Context(), tt.domainTransaction)
			if err != nil && !tt.wantErr {
				t.Errorf("Add() failed: %v", err)
				return
			}
		})
	}
}
