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

func TestOrderList(t *testing.T) {
	type m struct {
		userGetter  *mocks.UserGetter
		orderGetter *mocks.OrderAllByUserGetter
	}

	type args struct {
		ctx    context.Context
		userID uuid.UUID
	}

	userUUID := uuid.New()
	uploadedAt := time.Now()
	processedAt := time.Now()
	ordersMock := []*entity.Order{
		{
			Number:      entity.OrderNumber("12345678903"),
			UserID:      userUUID,
			Status:      entity.New,
			Accrual:     111.1,
			UploadedAt:  &uploadedAt,
			ProcessedAt: &processedAt,
		},
		{
			Number:      entity.OrderNumber("3272700463"),
			UserID:      userUUID,
			Status:      entity.New,
			Accrual:     222.2,
			UploadedAt:  &uploadedAt,
			ProcessedAt: &processedAt,
		},
	}

	tests := []struct {
		name        string
		args        *args
		want        []*entity.Order
		wantErr     error
		mockPrepare func(a *args) *m
	}{
		{
			"success",
			&args{
				context.Background(),
				userUUID,
			},
			ordersMock,
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

				orderGetter := mocks.NewOrderAllByUserGetter(t)
				orderGetter.EXPECT().
					GetAllByUser(a.ctx, mock.AnythingOfType("*entity.User")).
					Return(ordersMock, nil)

				return &m{userGetter, orderGetter}
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

				orderGetter := mocks.NewOrderAllByUserGetter(t)

				return &m{userGetter, orderGetter}
			},
		},
		{
			"fail with order not found error",
			&args{
				context.Background(),
				uuid.New(),
			},
			nil,
			ErrOrderNotFound,
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

				orderGetter := mocks.NewOrderAllByUserGetter(t)
				orderGetter.EXPECT().
					GetAllByUser(a.ctx, mock.AnythingOfType("*entity.User")).
					Return(nil, ErrOrderNotFound)

				return &m{userGetter, orderGetter}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			withdraws, err := OrderList(
				tt.args.ctx,
				tt.args.userID.String(),
				m.orderGetter,
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
