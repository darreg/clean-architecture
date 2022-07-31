package usecase

import (
	"context"
	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/alrund/yp-1-project/mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestWithdrawList(t *testing.T) {
	type mock struct {
		userGetter     *mocks.UserGetter
		withdrawGetter *mocks.WithdrawAllByUserGetter
	}

	type args struct {
		ctx    context.Context
		userID uuid.UUID
		user   *entity.User
	}

	userUUID := uuid.New()
	userMock := &entity.User{
		ID:           userUUID,
		Login:        "User login",
		PasswordHash: "Password hash",
		Current:      555.55,
	}
	processedAt := time.Now()
	withdrawsMock := []*entity.Withdraw{
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
	mockPrepare := func(m *mock, ctx context.Context, userID uuid.UUID, user *entity.User) {
		userGetterMock := m.userGetter.On("Get", ctx, userID)
		if userID == userMock.ID {
			userGetterMock.Return(userMock, nil)
		} else {
			userGetterMock.Return(nil, ErrUserNotFound)
		}

		withdrawGetter := m.withdrawGetter.On("GetAllByUser", ctx, user)
		if user.ID == userMock.ID {
			withdrawGetter.Return(withdrawsMock, nil)
		} else {
			withdrawGetter.Return(nil, ErrWithdrawNotFound)
		}
	}

	tests := []struct {
		name        string
		args        *args
		want        []*entity.Withdraw
		wantErr     error
		mockPrepare func(m *mock, ctx context.Context, userID uuid.UUID, user *entity.User)
	}{
		{
			"success",
			&args{
				context.Background(),
				userUUID,
				userMock,
			},
			withdrawsMock,
			nil,
			mockPrepare,
		},
		{
			"fail",
			&args{
				context.Background(),
				uuid.New(),
				userMock,
			},
			nil,
			ErrUserNotFound,
			mockPrepare,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mock{
				userGetter:     new(mocks.UserGetter),
				withdrawGetter: new(mocks.WithdrawAllByUserGetter),
			}

			if tt.mockPrepare != nil {
				tt.mockPrepare(m, tt.args.ctx, tt.args.userID, tt.args.user)
			}

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
