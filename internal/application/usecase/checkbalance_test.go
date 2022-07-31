package usecase

import (
	"context"
	"testing"

	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type UserGetterMock struct {
	getFn func(ctx context.Context, userID uuid.UUID) (*entity.User, error)
}

func (ug *UserGetterMock) Get(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	return ug.getFn(ctx, userID)
}

type WithdrawnerMock struct {
	getWithdrawnFn func(ctx context.Context, user *entity.User) (float32, error)
}

func (w *WithdrawnerMock) GetWithdrawn(ctx context.Context, user *entity.User) (float32, error) {
	return w.getWithdrawnFn(ctx, user)
}

func TestCheckBalance(t *testing.T) {
	userUUID := uuid.New()
	tests := []struct {
		name            string
		userID          string
		userGetterMock  *UserGetterMock
		withdrawnerMock *WithdrawnerMock
		balance         *Balance
		err             error
	}{
		{
			"success",
			userUUID.String(),
			&UserGetterMock{
				getFn: func(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
					if userID != userUUID {
						return nil, ErrUserNotFound
					}
					return &entity.User{
						ID:           userUUID,
						Login:        "User login",
						PasswordHash: "Password hash",
						Current:      555.55,
					}, nil
				},
			},
			&WithdrawnerMock{
				getWithdrawnFn: func(ctx context.Context, user *entity.User) (float32, error) {
					return 666.66, nil
				},
			},
			&Balance{
				Current:   555.55,
				Withdrawn: 666.66,
			},
			nil,
		},
		{
			"fail",
			uuid.NewString(),
			&UserGetterMock{
				getFn: func(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
					if userID != userUUID {
						return nil, ErrUserNotFound
					}
					return &entity.User{
						ID:           userUUID,
						Login:        "User login",
						PasswordHash: "Password hash",
						Current:      555.55,
					}, nil
				},
			},
			&WithdrawnerMock{
				getWithdrawnFn: func(ctx context.Context, user *entity.User) (float32, error) {
					return 666.66, nil
				},
			},
			&Balance{
				Current:   555.55,
				Withdrawn: 666.66,
			},
			ErrUserNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			balance, err := CheckBalance(
				context.Background(),
				tt.userID,
				tt.userGetterMock,
				tt.withdrawnerMock,
			)
			if tt.err != nil {
				assert.Error(t, tt.err, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.balance.Current, balance.Current)
			assert.Equal(t, tt.balance.Withdrawn, balance.Withdrawn)
		})
	}
}
