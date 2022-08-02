//go:build unit

package usecase

import (
	"context"
	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/alrund/yp-1-project/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAccrualProcess(t *testing.T) {
	type m struct {
		userAccrualer *mocks.UserAccrualer
		orderChanger  *mocks.OrderWithCheckChanger
		transactor    *mocks.Transactor
	}

	type args struct {
		ctx           context.Context
		user          *entity.User
		accrualResult *AccrualResult
	}

	transactor := mocks.NewTransactor(t)
	transactor.EXPECT().
		WithinTransaction(context.Background(), mock.Anything).
		Call.
		Return(
			func(ctx context.Context, tFunc func(ctx context.Context) error) error {
				return tFunc(ctx)
			},
		)

	tests := []struct {
		name        string
		args        *args
		wantErr     error
		mockPrepare func(a *args) *m
	}{
		{
			"success",
			&args{
				context.Background(),
				&entity.User{
					ID: uuid.New(),
				},
				&AccrualResult{
					OrderNumber: "3272700463",
					Status:      entity.Processed,
					Accrual:     100,
				},
			},
			nil,
			func(a *args) *m {
				userAccrualer := mocks.NewUserAccrualer(t)
				userAccrualer.EXPECT().
					Accrual(a.ctx, a.user, a.accrualResult.Accrual).
					Return(nil).
					Once()

				orderChanger := mocks.NewOrderWithCheckChanger(t)
				orderChanger.EXPECT().
					Get(a.ctx, a.accrualResult.OrderNumber).
					Return(&entity.Order{}, nil)
				orderChanger.EXPECT().
					Change(a.ctx, mock.AnythingOfType("*entity.Order")).
					Return(nil)

				return &m{userAccrualer, orderChanger, transactor}
			},
		},
		{
			"success with order not Processed status",
			&args{
				context.Background(),
				&entity.User{
					ID: uuid.New(),
				},
				&AccrualResult{
					OrderNumber: "3272700463",
					Status:      entity.New,
					Accrual:     100,
				},
			},
			nil,
			func(a *args) *m {
				userAccrualer := mocks.NewUserAccrualer(t)

				orderChanger := mocks.NewOrderWithCheckChanger(t)
				orderChanger.EXPECT().
					Get(a.ctx, a.accrualResult.OrderNumber).
					Return(&entity.Order{}, nil)
				orderChanger.EXPECT().
					Change(a.ctx, mock.AnythingOfType("*entity.Order")).
					Return(nil)

				return &m{userAccrualer, orderChanger, transactor}
			},
		},
		{
			"success with zero accrual status",
			&args{
				context.Background(),
				&entity.User{
					ID: uuid.New(),
				},
				&AccrualResult{
					OrderNumber: "3272700463",
					Status:      entity.New,
					Accrual:     0,
				},
			},
			nil,
			func(a *args) *m {
				userAccrualer := mocks.NewUserAccrualer(t)

				orderChanger := mocks.NewOrderWithCheckChanger(t)
				orderChanger.EXPECT().
					Get(a.ctx, a.accrualResult.OrderNumber).
					Return(&entity.Order{}, nil)
				orderChanger.EXPECT().
					Change(a.ctx, mock.AnythingOfType("*entity.Order")).
					Return(nil)

				return &m{userAccrualer, orderChanger, transactor}
			},
		},
		{
			"fail with order not found",
			&args{
				context.Background(),
				&entity.User{
					ID: uuid.New(),
				},
				&AccrualResult{
					OrderNumber: "3272700463",
					Status:      entity.Processed,
					Accrual:     100,
				},
			},
			ErrOrderNotFound,
			func(a *args) *m {
				userAccrualer := mocks.NewUserAccrualer(t)

				orderChanger := mocks.NewOrderWithCheckChanger(t)
				orderChanger.EXPECT().
					Get(a.ctx, a.accrualResult.OrderNumber).
					Return(nil, ErrOrderNotFound)

				return &m{userAccrualer, orderChanger, transactor}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			err := AccrualProcess(
				tt.args.ctx,
				tt.args.accrualResult,
				tt.args.user,
				m.userAccrualer,
				m.orderChanger,
				m.transactor,
			)

			if tt.wantErr != nil {
				assert.Error(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
