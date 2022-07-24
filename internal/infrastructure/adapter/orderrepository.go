package adapter

import (
	"github.com/alrund/yp-1-project/internal/domain/entity"
)

type OrderRepository struct{}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{}
}

func (o OrderRepository) Find(number string) *entity.Order {
	// TODO implement me
	panic("implement me")
}

func (o OrderRepository) FindByUser(user entity.User) *entity.Order {
	// TODO implement me
	panic("implement me")
}
