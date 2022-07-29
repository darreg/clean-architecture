package port

import (
	"context"
	"database/sql"

	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/google/uuid"
)

type Transactor interface {
	WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error
	QueryContext(ctx context.Context, db *sql.DB, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, db *sql.DB, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, db *sql.DB, query string, args ...interface{}) (sql.Result, error)
}

type UserRepository interface {
	Transactor
	Get(ctx context.Context, userID uuid.UUID) (*entity.User, error)
	GetByLogin(ctx context.Context, login string) (*entity.User, error)
	GetByCredential(ctx context.Context, login, passwordHash string) (*entity.User, error)
	Add(ctx context.Context, user *entity.User) error
	Withdraw(ctx context.Context, user *entity.User, sum int) error
	Change(ctx context.Context, user *entity.User) error
	ChangePassword(ctx context.Context, user *entity.User) error
	Remove(ctx context.Context, userID uuid.UUID) error
}

type OrderRepository interface {
	Transactor
	Get(ctx context.Context, number string) (*entity.Order, error)
	GetAllByUser(ctx context.Context, user *entity.User) ([]*entity.Order, error)
	Add(ctx context.Context, order *entity.Order) error
	Change(ctx context.Context, order *entity.Order) error
	Remove(ctx context.Context, number string) error
}

type WithdrawRepository interface {
	Transactor
	Get(ctx context.Context, withdrawID uuid.UUID) (*entity.Withdraw, error)
	GetWithdrawn(ctx context.Context, user *entity.User) (int, error)
	GetAllByUser(ctx context.Context, user *entity.User) ([]*entity.Withdraw, error)
	Add(ctx context.Context, order *entity.Withdraw) error
	Change(ctx context.Context, order *entity.Withdraw) error
	Remove(ctx context.Context, withdrawID uuid.UUID) error
}
