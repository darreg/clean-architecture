package port

import (
	"context"

	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/google/uuid"
)

type WithdrawGetter interface {
	Get(ctx context.Context, withdrawID uuid.UUID) (*entity.Withdraw, error)
}

type WithdrawAllByUserGetter interface {
	GetAllByUser(ctx context.Context, user *entity.User) ([]*entity.Withdraw, error)
}

type Withdrawner interface {
	GetWithdrawn(ctx context.Context, user *entity.User) (float32, error)
}

type WithdrawAdder interface {
	Add(ctx context.Context, order *entity.Withdraw) error
}

type WithdrawChanger interface {
	Change(ctx context.Context, order *entity.Withdraw) error
}

type WithdrawRemover interface {
	Remove(ctx context.Context, withdrawID uuid.UUID) error
}

type WithdrawRepository interface {
	WithdrawGetter
	WithdrawAllByUserGetter
	Withdrawner
	WithdrawAdder
	WithdrawChanger
	WithdrawRemover
}
