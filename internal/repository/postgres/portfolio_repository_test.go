package postgres_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/gorkagg10/equity-calculator-api/internal/domain"
	"github.com/gorkagg10/equity-calculator-api/internal/repository/postgres"
	"github.com/stretchr/testify/require"
)

func TestPortfolioRepository_AddPortfolio(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		pgClient *sql.DB
		// Named input parameters for target function.
		portfolio *domain.Portfolio
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		db := startDatabase(t)

		ctx := context.Background()
		require.NoError(t, postgres.Migrate(ctx, db))

		t.Run(tt.name, func(t *testing.T) {
			p := postgres.NewPortfolioRepository(tt.pgClient)
			gotErr := p.AddPortfolio(t.Context(), tt.portfolio)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("AddPortfolio() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("AddPortfolio() succeeded unexpectedly")
			}
		})
	}
}
