package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/alrund/yp-1-project/internal/application/usecase"
	"github.com/alrund/yp-1-project/internal/domain/entity"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (o OrderRepository) Get(ctx context.Context, number string) (*entity.Order, error) {
	var order entity.Order
	var uploadedAt time.Time
	var processedAt sql.NullTime

	err := o.db.QueryRowContext(ctx,
		"SELECT number, user_id, status, accrual, uploaded_at, processed_at FROM orders WHERE number = $1", number,
	).Scan(&order.Number, &order.UserID, &order.Status, &order.Accrual, &uploadedAt, &processedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrOrderNotFound
		}
		return nil, err
	}

	order.UploadedAt = &uploadedAt

	if processedAt.Valid {
		order.ProcessedAt = &processedAt.Time
	}

	return &order, nil
}

func (o OrderRepository) GetAllByUser(ctx context.Context, user *entity.User) ([]*entity.Order, error) {
	rows, err := o.db.QueryContext(ctx,
		"SELECT number, user_id, status, accrual, uploaded_at, processed_at FROM orders WHERE user_id = $1", user.ID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrOrderNotFound
		}
		return nil, err
	}

	defer rows.Close()

	orders := make([]*entity.Order, 0)
	for rows.Next() {
		var order entity.Order
		var uploadedAt time.Time
		var processedAt sql.NullTime
		err = rows.Scan(&order.Number, &order.UserID, &order.Status, &order.Accrual, &uploadedAt, &processedAt)
		if err != nil {
			return nil, err
		}

		order.UploadedAt = &uploadedAt

		if processedAt.Valid {
			order.ProcessedAt = &processedAt.Time
		}

		orders = append(orders, &order)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (o OrderRepository) Add(ctx context.Context, order *entity.Order) error {
	_, err := o.db.ExecContext(ctx,
		"INSERT INTO orders(number, user_id, status, accrual, uploaded_at, processed_at) VALUES($1, $2, $3, $4, $5, $6)",
		order.Number, order.UserID, order.Status, order.Accrual, order.UploadedAt, order.ProcessedAt)
	if err != nil {
		return err
	}

	return nil
}

func (o OrderRepository) Change(ctx context.Context, order *entity.Order) error {
	_, err := o.db.ExecContext(ctx,
		"UPDATE orders SET user_id=$2, status=$3, accrual=$4, processed_at=$5 WHERE number=$1",
		order.Number, order.UserID, order.Status, order.Accrual, order.ProcessedAt)
	if err != nil {
		return err
	}

	return nil
}

func (o OrderRepository) Remove(ctx context.Context, number string) error {
	_, err := o.db.ExecContext(ctx, "DELETE FROM orders WHERE number=$1", number)
	return err
}
