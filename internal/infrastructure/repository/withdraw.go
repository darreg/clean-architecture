package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/alrund/yp-1-project/internal/application/usecase"
	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/alrund/yp-1-project/internal/infrastructure/adapter"
	"github.com/google/uuid"
)

type WithdrawRepository struct {
	tx *adapter.Transactor
	db *sql.DB
}

func NewWithdrawRepository(tx *adapter.Transactor, db *sql.DB) *WithdrawRepository {
	return &WithdrawRepository{tx: tx, db: db}
}

func (w WithdrawRepository) Get(ctx context.Context, withdrawID uuid.UUID) (*entity.Withdraw, error) {
	var withdraw entity.Withdraw
	var processedAt time.Time

	err := w.tx.QueryRowContext(ctx,
		"SELECT id, order_number, user_id, sum, processed_at FROM withdraws WHERE id = $1", withdrawID,
	).Scan(&withdraw.ID, &withdraw.OrderNumber, &withdraw.UserID, &withdraw.Sum, &processedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrWithdrawNotFound
		}
		return nil, err
	}

	withdraw.ProcessedAt = &processedAt

	return &withdraw, nil
}

func (w WithdrawRepository) GetWithdrawn(ctx context.Context, user *entity.User) (float32, error) {
	var withdrawn sql.NullFloat64

	err := w.tx.QueryRowContext(ctx,
		"SELECT SUM(sum) as withdrawn FROM withdraws WHERE user_id = $1", user.ID,
	).Scan(&withdrawn)
	if err != nil {
		return 0, err
	}

	if !withdrawn.Valid {
		return 0, nil
	}

	return float32(withdrawn.Float64), nil
}

func (w WithdrawRepository) GetAllByUser(ctx context.Context, user *entity.User) ([]entity.Withdraw, error) {
	rows, err := w.tx.QueryContext(ctx,
		"SELECT id, order_number, user_id, sum, processed_at FROM withdraws WHERE user_id = $1", user.ID,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrWithdrawNotFound
		}
		return nil, err
	}

	defer rows.Close()

	withdraws := make([]entity.Withdraw, 0)
	for rows.Next() {
		var withdraw entity.Withdraw
		var processedAt time.Time
		err = rows.Scan(&withdraw.ID, &withdraw.OrderNumber, &withdraw.UserID, &withdraw.Sum, &processedAt)
		if err != nil {
			return nil, err
		}

		withdraw.ProcessedAt = &processedAt

		withdraws = append(withdraws, withdraw)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if len(withdraws) == 0 {
		return nil, usecase.ErrWithdrawNotFound
	}

	return withdraws, nil
}

func (w WithdrawRepository) Add(ctx context.Context, withdraw *entity.Withdraw) error {
	_, err := w.tx.ExecContext(ctx,
		"INSERT INTO withdraws(id, order_number, user_id, sum, processed_at) VALUES($1, $2, $3, $4, $5)",
		withdraw.ID, withdraw.OrderNumber.String(), withdraw.UserID, withdraw.Sum, withdraw.ProcessedAt)
	if err != nil {
		return err
	}

	return nil
}

func (w WithdrawRepository) Change(ctx context.Context, withdraw *entity.Withdraw) error {
	_, err := w.tx.ExecContext(ctx,
		"UPDATE withdraws SET order_number=$2, user_id=$3, sum=$4, processed_at=$5 WHERE id=$1",
		withdraw.ID, withdraw.OrderNumber.String(), withdraw.UserID, withdraw.Sum, withdraw.ProcessedAt)
	if err != nil {
		return err
	}

	return nil
}

func (w WithdrawRepository) Remove(ctx context.Context, withdrawID uuid.UUID) error {
	_, err := w.tx.ExecContext(ctx, "DELETE FROM withdraws WHERE id=$1", withdrawID)
	return err
}
