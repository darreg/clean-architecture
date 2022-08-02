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

func TestAddWithdraw(t *testing.T) {
	type m struct {
		userAddWithdrawer *mocks.UserAddWithdrawer
		withdrawAdder     *mocks.WithdrawAdder
		transactor        *mocks.Transactor
	}

	type args struct {
		ctx    context.Context
		number string
		sum    float32
		userID uuid.UUID
	}

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
				"3272700463",
				100.0,
				uuid.New(),
			},
			nil,
			func(a *args) *m {
				userAddWithdrawer := mocks.NewUserAddWithdrawer(t)
				userAddWithdrawer.EXPECT().
					Get(a.ctx, a.userID).
					Return(&entity.User{
						ID:           a.userID,
						Login:        "User login",
						PasswordHash: "Password hash",
						Current:      200,
					}, nil)
				userAddWithdrawer.EXPECT().
					Withdraw(a.ctx, mock.AnythingOfType("*entity.User"), a.sum).
					Return(nil).
					Once()

				withdrawAdder := mocks.NewWithdrawAdder(t)
				withdrawAdder.EXPECT().
					Add(a.ctx, mock.AnythingOfType("*entity.Withdraw")).
					Return(nil).
					Once()

				transactor := mocks.NewTransactor(t)
				transactor.EXPECT().
					WithinTransaction(context.Background(), mock.Anything).
					Call.
					Return(
						func(ctx context.Context, tFunc func(ctx context.Context) error) error {
							return tFunc(ctx)
						},
					)

				return &m{userAddWithdrawer, withdrawAdder, transactor}
			},
		},
		{
			"fail with not enough funds",
			&args{
				context.Background(),
				"3272700463",
				100.0,
				uuid.New(),
			},
			ErrNotEnoughFunds,
			func(a *args) *m {
				userAddWithdrawer := mocks.NewUserAddWithdrawer(t)
				userAddWithdrawer.EXPECT().
					Get(a.ctx, a.userID).
					Return(&entity.User{
						ID:           a.userID,
						Login:        "User login",
						PasswordHash: "Password hash",
						Current:      0,
					}, nil)

				withdrawAdder := mocks.NewWithdrawAdder(t)

				transactor := mocks.NewTransactor(t)

				return &m{userAddWithdrawer, withdrawAdder, transactor}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			err := AddWithdraw(
				tt.args.ctx,
				tt.args.number,
				tt.args.sum,
				tt.args.userID.String(),
				m.userAddWithdrawer,
				m.withdrawAdder,
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
