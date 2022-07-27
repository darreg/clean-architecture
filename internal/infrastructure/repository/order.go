package repository

import (
	"database/sql"

	"github.com/alrund/yp-1-project/internal/domain/entity"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (o OrderRepository) Get(number string) (*entity.Order, error) {
	// TODO implement me
	panic("implement me")
}

func (o OrderRepository) GetByUser(user entity.User) (*entity.Order, error) {
	// TODO implement me
	panic("implement me")
}
