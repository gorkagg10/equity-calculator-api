package postgres_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gorkagg10/equity-calculator-api/internal/domain"
	"github.com/gorkagg10/equity-calculator-api/internal/repository/postgres"
	"github.com/stretchr/testify/require"
)

func TestAssetRepository_FindByID(t *testing.T) {
	assetID := uuid.New()
	timeNow := time.Now().UTC()

	exampleAsset, err := domain.NewAsset(
		assetID,
		domain.NewAssetData(
			"TST",
			"test",
			10.00,
			"EUR",
			"ESP",
		),
		timeNow,
		timeNow,
	)
	require.NoError(t, err)

	tests := []struct {
		name    string // description of this test case
		before  func(t *testing.T, assetRepository *postgres.AssetRepository)
		assetID uuid.UUID
		want    *domain.Asset
		wantErr bool
	}{
		{
			name:    "asset not found",
			assetID: assetID,
			wantErr: true,
		},
		{
			name: "success finding asset by ID",
			before: func(t *testing.T, assetRepository *postgres.AssetRepository) {
				require.NoError(t, assetRepository.Add(t.Context(), exampleAsset))
			},
			assetID: assetID,
			want:    exampleAsset,
		},
	}
	for _, tt := range tests {
		db := startDatabase(t)
		require.NotNil(t, db)
		require.NoError(t, postgres.Migrate(db, "app-test", "migrations"))

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			a := postgres.NewAssetRepository(db)
			if tt.before != nil {
				tt.before(t, a)
			}
			got, err := a.FindByID(t.Context(), tt.assetID)
			if err != nil && !tt.wantErr {
				t.Errorf("FindByID() failed: %v", err)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}
