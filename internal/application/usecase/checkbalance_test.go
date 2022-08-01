package usecase

import (
	"context"
	"testing"

	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/alrund/yp-1-project/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCheckBalance(t *testing.T) {
	type m struct {
		userGetter  *mocks.UserGetter
		withdrawner *mocks.Withdrawner
	}

	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}

	const Current float32 = 555.55
	const Withdrawn float32 = 555.55

	userUUID := uuid.New()

	tests := []struct {
		name        string
		args        *args
		want        *Balance
		wantErr     error
		mockPrepare func(a *args) *m
	}{
		{
			"success",
			&args{
				context.Background(),
				userUUID,
			},
			&Balance{Current: Current, Withdrawn: Withdrawn},
			nil,
			func(a *args) *m {
				userGetter := mocks.NewUserGetter(t)
				userGetter.EXPECT().
					Get(a.ctx, a.userID).
					Return(&entity.User{
						ID:           a.userID,
						Login:        "User login",
						PasswordHash: "Password hash",
						Current:      Current,
					}, nil)

				withdrawner := mocks.NewWithdrawner(t)
				withdrawner.EXPECT().
					GetWithdrawn(a.ctx, mock.AnythingOfType("*entity.User")).
					Return(Withdrawn, nil)

				return &m{userGetter, withdrawner}
			},
		},
		{
			"fail",
			&args{
				context.Background(),
				userUUID,
			},
			nil,
			ErrUserNotFound,
			func(a *args) *m {
				userGetter := mocks.NewUserGetter(t)
				userGetter.EXPECT().
					Get(a.ctx, a.userID).
					Return(nil, ErrUserNotFound)

				withdrawner := mocks.NewWithdrawner(t)

				return &m{userGetter, withdrawner}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			balance, err := CheckBalance(
				tt.args.ctx,
				tt.args.userID.String(),
				m.userGetter,
				m.withdrawner,
			)

			if tt.wantErr != nil {
				assert.Error(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want.Current, balance.Current)
			assert.Equal(t, tt.want.Withdrawn, balance.Withdrawn)
		})
	}
}
