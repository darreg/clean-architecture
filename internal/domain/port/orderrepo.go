package port

import (
	"context"

	"github.com/alrund/yp-1-project/internal/domain/entity"
)

type OrderGetter interface {
	Get(ctx context.Context, number string) (*entity.Order, error)
}

type OrderAllByUserGetter interface {
	GetAllByUser(ctx context.Context, user *entity.User) ([]*entity.Order, error)
}

type OrderAdder interface {
	Add(ctx context.Context, order *entity.Order) error
}

type OrderChanger interface {
	Change(ctx context.Context, order *entity.Order) error
}

type OrderRemover interface {
	Remove(ctx context.Context, number string) error
}

type OrderRepository interface {
	Transactor
	OrderGetter
	OrderAllByUserGetter
	OrderAdder
	OrderChanger
	OrderRemover
}
