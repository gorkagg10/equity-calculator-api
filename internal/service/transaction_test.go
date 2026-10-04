package service_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/gorkagg10/equity-calculator-api/internal/domain"
	"github.com/gorkagg10/equity-calculator-api/internal/service"
)

func TestTransaction_Add(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		portfolioID     uuid.UUID
		assetID         uuid.UUID
		unitPrice       float64
		quantity        float64
		currency        string
		transactionType string
	}
	tests := []struct {
		name    string
		before  func(ctrl *gomock.Controller) *service.Transaction
		args    args
		want    *domain.Transaction
		wantErr bool
	}{
		{
			name: "portfolio not found",
			before: func(ctrl *gomock.Controller) *service.Transaction {
				portfolioRepository := domain.NewMockPortfolioRepository(ctrl)
				portfolioRepository.EXPECT().FindByID(gomock.Any(), gomock.Any()).Return(nil, errors.New("portfolio not found"))

				return service.NewTransaction(
					nil,
					portfolioRepository,
					nil,
					nil,
				)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tr := tt.before(ctrl)
			got, err := tr.Add(t.Context(), tt.args.portfolioID, tt.args.assetID, tt.args.unitPrice, tt.args.quantity, tt.args.currency, tt.args.transactionType)
			if (err != nil) && !tt.wantErr {
				t.Errorf("Add() failed: %v", err)
				return
			}
			require.Equal(t, tt.want, got)
		})
	}
}
