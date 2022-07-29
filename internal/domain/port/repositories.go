package port

import (
	"context"
	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/google/uuid"
)

type UserRepository interface {
	Get(ctx context.Context, userID uuid.UUID) (*entity.User, error)
	GetByLogin(ctx context.Context, login string) (*entity.User, error)
	GetByCredential(ctx context.Context, login, passwordHash string) (*entity.User, error)
	Add(ctx context.Context, user *entity.User) error
	Change(ctx context.Context, user *entity.User) error
	ChangePassword(ctx context.Context, user *entity.User) error
	Remove(ctx context.Context, userID uuid.UUID) error
}

type OrderRepository interface {
	Get(ctx context.Context, number string) (*entity.Order, error)
	GetAllByUser(ctx context.Context, user *entity.User) ([]*entity.Order, error)
	Add(ctx context.Context, order *entity.Order) error
	Change(ctx context.Context, order *entity.Order) error
	Remove(ctx context.Context, number string) error
}

type WithdrawRepository interface {
	Get(ctx context.Context, withdrawID uuid.UUID) (*entity.Withdraw, error)
	GetAllByUser(ctx context.Context, user *entity.User) ([]*entity.Withdraw, error)
	Add(ctx context.Context, order *entity.Withdraw) error
	Change(ctx context.Context, order *entity.Withdraw) error
	Remove(ctx context.Context, withdrawID uuid.UUID) error
}
