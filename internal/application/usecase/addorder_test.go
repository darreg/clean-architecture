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
	"time"
)

func TestAddOrder(t *testing.T) {
	type m struct {
		userGetter *mocks.UserGetter
		orderAdder *mocks.OrderWithCheckAdder
	}

	type args struct {
		ctx    context.Context
		number string
		userID uuid.UUID
	}

	userUUID := uuid.New()

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
				"5251743000",
				userUUID,
			},
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

				orderAdder := mocks.NewOrderWithCheckAdder(t)
				orderAdder.EXPECT().
					Get(a.ctx, mock.AnythingOfType("string")).
					Return(nil, ErrOrderNotFound)
				orderAdder.EXPECT().
					Add(a.ctx, mock.AnythingOfType("*entity.Order")).
					Return(nil).
					Once()
				return &m{userGetter, orderAdder}
			},
		},
		{
			"fail with order already uploaded",
			&args{
				context.Background(),
				"5251743000",
				userUUID,
			},
			ErrOrderAlreadyUploaded,
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

				uploadedAt := time.Now()
				processedAt := time.Now()
				orderAdder := mocks.NewOrderWithCheckAdder(t)
				orderAdder.EXPECT().
					Get(a.ctx, mock.AnythingOfType("string")).
					Return(&entity.Order{
						Number:      "5251743000",
						UserID:      userUUID,
						Status:      entity.New,
						Accrual:     111.1,
						UploadedAt:  &uploadedAt,
						ProcessedAt: &processedAt,
					}, nil)

				return &m{userGetter, orderAdder}
			},
		},
		{
			"fail with order already uploaded by another user",
			&args{
				context.Background(),
				"5251743000",
				userUUID,
			},
			ErrOrderAlreadyUploadedAnotherUser,
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

				uploadedAt := time.Now()
				processedAt := time.Now()
				orderAdder := mocks.NewOrderWithCheckAdder(t)
				orderAdder.EXPECT().
					Get(a.ctx, mock.AnythingOfType("string")).
					Return(&entity.Order{
						Number:      "5251743000",
						UserID:      uuid.New(),
						Status:      entity.New,
						Accrual:     111.1,
						UploadedAt:  &uploadedAt,
						ProcessedAt: &processedAt,
					}, nil)

				return &m{userGetter, orderAdder}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.mockPrepare(tt.args)

			err := AddOrder(
				tt.args.ctx,
				tt.args.number,
				tt.args.userID.String(),
				m.orderAdder,
				m.userGetter,
			)

			if tt.wantErr != nil {
				assert.Error(t, tt.wantErr, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
