//go:build unit

package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/alrund/yp-1-project/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWithdrawList(t *testing.T) {
	type m struct {
		userGetter     *mocks.UserGetter
		withdrawGetter *mocks.WithdrawAllByUserGetter
	}

	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}

	userUUID := uuid.New()
	processedAt := time.Now()
	withdrawsMock := []entity.Withdraw{
		{
			ID:          uuid.New(),
			UserID:      userUUID,
			OrderNumber: entity.OrderNumber("xxx"),
			Sum:         100.1,
			ProcessedAt: &processedAt,
		},
		{
			ID:          uuid.New(),
			UserID:      userUUID,
			OrderNumber: entity.OrderNumber("yyy"),
			Sum:         200.2,
			ProcessedAt: &processedAt,
		},
	}

	tests := []struct {
		name        string
		args        *args
		want        []entity.Withdraw
		wantErr     error
		mockPrepare func(a *args) *m
	}{
		{
			"success",
			&args{
				context.Background(),
				userUUID,
			},
			withdrawsMock,
			nil,
			func(a *args) *m {
				userGetter := mocks.NewUserGetter(t)
				userGetter.EXPECT().
					Get(a.ctx, a.userID).
					Return(&entity.User{
						ID:           a.userID,
						Login:        "User login",
						PasswordHash: "Password hash",
						Current:      0,
					}, nil)

				withdrawGetter := mocks.NewWithdrawAllByUserGetter(t)
				withdrawGetter.EXPECT().
					GetAllByUser(a.ctx, mock.AnythingOfType("*entity.User")).
					Return(withdrawsMock, nil)

				return &m{userGetter, withdrawGetter}
			},
		},
		{
			"fail with user not found error",
			&args{
				context.Background(),
				uuid.New(),
			},
			nil,
			ErrUserNotFound,
			func(a *args) *m {
				userGetter := mocks.NewUserGetter(t)
				userGetter.EXPECT().
					Get(a.ctx, a.userID).
					Return(nil, ErrUserNotFound)

				withdrawGetter := mocks.NewWithdrawAllByUserGetter(t)

				return &m{userGetter, withdrawGetter}
			},
		},
		{
			"fail with withdraw not found error",
			&args{
				context.Background(),
				uuid.New(),
			},
			nil,
			ErrWithdrawNotFound,
			func(a *args) *m {
				userGetter := mocks.NewUserGetter(t)
				userGetter.EXPECT().
					Get(a.ctx, a.userID).
					Return(&entity.User{
						ID:           a.userID,
						Login:        "User login",
						PasswordHash: "Password hash",
						Current:      0,
					}, nil)

				withdrawGetter := mocks.NewWithdrawAllByUserGetter(t)
				withdrawGetter.EXPECT().
					GetAllByUser(a.ctx, mock.AnythingOfType("*entity.User")).
					Return(nil, ErrWithdrawNotFound)

				return &m{userGetter, withdrawGetter}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			withdraws, err := WithdrawList(
				tt.args.ctx,
				tt.args.userID.String(),
				m.withdrawGetter,
				m.userGetter,
			)

			if tt.wantErr != nil {
				assert.Error(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, withdraws)
		})
	}
}
