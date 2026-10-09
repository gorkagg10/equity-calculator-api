package postgres_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gorkagg10/equity-calculator-api/internal/domain"
	"github.com/gorkagg10/equity-calculator-api/internal/repository/postgres"
	"github.com/stretchr/testify/require"
)

func TestPortfolioRepository_AddPortfolio(t *testing.T) {
	examplePortfolio, err := domain.NewPortfolio(
		uuid.New(),
		"test",
		time.Now(),
		time.Now(),
	)
	require.NoError(t, err)

	tests := []struct {
		name      string
		before    func(t *testing.T, portfolioRepository *postgres.PortfolioRepository)
		portfolio *domain.Portfolio
		wantErr   bool
	}{
		{
			name: "error - user already exists",
			before: func(t *testing.T, portfolioRepository *postgres.PortfolioRepository) {
				require.NoError(t, portfolioRepository.AddPortfolio(t.Context(), examplePortfolio))
			},
			portfolio: examplePortfolio,
			wantErr:   true,
		},
		{
			name:      "success",
			portfolio: examplePortfolio,
		},
	}
	for _, tt := range tests {
		db := startDatabase(t)
		require.NotNil(t, db)
		require.NoError(t, postgres.Migrate(db, "app-test", "migrations"))

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := postgres.NewPortfolioRepository(db)
			if tt.before != nil {
				tt.before(t, p)
			}
			gotErr := p.AddPortfolio(t.Context(), tt.portfolio)
			if gotErr != nil && !tt.wantErr {
				t.Errorf("AddPortfolio() failed: %v", gotErr)
				return
			}
		})
	}
}

func TestPortfolioRepository_FindByID(t *testing.T) {
	portfolioID := uuid.New()

	examplePortfolio, err := domain.NewPortfolio(
		portfolioID,
		"test",
		time.Now().UTC(),
		time.Now().UTC(),
	)
	require.NoError(t, err)

	tests := []struct {
		name        string
		before      func(t *testing.T, portfolioRepository *postgres.PortfolioRepository)
		portfolioID uuid.UUID
		want        *domain.Portfolio
		wantErr     bool
	}{
		{
			name:        "portfolio not found",
			portfolioID: portfolioID,
			wantErr:     true,
		},
		{
			name: "success finding portfolio by ID",
			before: func(t *testing.T, portfolioRepository *postgres.PortfolioRepository) {
				require.NoError(t, portfolioRepository.AddPortfolio(t.Context(), examplePortfolio))
			},
			portfolioID: portfolioID,
			want:        examplePortfolio,
		},
	}
	for _, tt := range tests {
		db := startDatabase(t)
		require.NotNil(t, db)
		require.NoError(t, postgres.Migrate(db, "app-test", "migrations"))

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := postgres.NewPortfolioRepository(db)
			if tt.before != nil {
				tt.before(t, p)
			}
			got, err := p.FindByID(t.Context(), tt.portfolioID)
			if err != nil && !tt.wantErr {
				t.Errorf("FindByID() failed: %v", err)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}
