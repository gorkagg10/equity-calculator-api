package postgres_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gorkagg10/equity-calculator-api/internal/domain"
	"github.com/gorkagg10/equity-calculator-api/internal/repository/postgres"
	"github.com/stretchr/testify/require"
)

func TestPositionRepository_LoadPosition(t *testing.T) {
	positionID := uuid.New()
	portfolioID := uuid.New()
	assetID := uuid.New()
	assetTime := time.Now().UTC()
	portfolioTime := time.Now().UTC()
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

	examplePosition := domain.NewPosition(
		positionID,
		portfolioID,
		assetID,
		10.10,
		"EUR",
		now,
		now,
	)

	type args struct {
		portfolioID uuid.UUID
		assetID     uuid.UUID
	}

	tests := []struct {
		name    string
		args    args
		before  func(t *testing.T, db *sql.DB, positionRepository *postgres.PositionRepository)
		want    *domain.Position
		wantErr bool
	}{
		{
			name: "portfolio missing error",
			args: args{
				portfolioID: portfolioID,
				assetID:     assetID,
			},
			wantErr: true,
		},
		{
			name: "asset missing error",
			args: args{
				portfolioID: portfolioID,
				assetID:     assetID,
			},
			before: func(t *testing.T, db *sql.DB, positionRepository *postgres.PositionRepository) {
				portfolioRepository := postgres.NewPortfolioRepository(db)
				require.NoError(t, portfolioRepository.AddPortfolio(t.Context(), examplePortfolio))
			},
			wantErr: true,
		},
		{
			name: "position not found",
			args: args{
				portfolioID: portfolioID,
				assetID:     assetID,
			},
			before: func(t *testing.T, db *sql.DB, positionRepository *postgres.PositionRepository) {
				portfolioRepository := postgres.NewPortfolioRepository(db)
				require.NoError(t, portfolioRepository.AddPortfolio(t.Context(), examplePortfolio))
				assetRepository := postgres.NewAssetRepository(db)
				require.NoError(t, assetRepository.Add(t.Context(), exampleAsset))
			},
			wantErr: true,
		},
		{
			name: "position successfully loaded",
			args: args{
				portfolioID: portfolioID,
				assetID:     assetID,
			},
			before: func(t *testing.T, db *sql.DB, positionRepository *postgres.PositionRepository) {
				portfolioRepository := postgres.NewPortfolioRepository(db)
				require.NoError(t, portfolioRepository.AddPortfolio(t.Context(), examplePortfolio))
				assetRepository := postgres.NewAssetRepository(db)
				require.NoError(t, assetRepository.Add(t.Context(), exampleAsset))
				require.NoError(t, positionRepository.Upsert(t.Context(), examplePosition))
			},
			want: examplePosition,
		},
	}
	for _, tt := range tests {
		db := startDatabase(t)
		require.NotNil(t, db)
		require.NoError(t, postgres.Migrate(db, "app-test", "migrations"))

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			p := postgres.NewPositionRepository(db)
			if tt.before != nil {
				tt.before(t, db, p)
			}
			got, err := p.LoadPosition(t.Context(), tt.args.portfolioID, tt.args.assetID)
			if err != nil && !tt.wantErr {
				t.Errorf("LoadPosition() failed: %v", err)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}
