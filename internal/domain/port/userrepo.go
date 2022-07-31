package port

import (
	"context"

	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/google/uuid"
)

type UserGetter interface {
	Get(ctx context.Context, userID uuid.UUID) (*entity.User, error)
}

type UserByLoginGetter interface {
	GetByLogin(ctx context.Context, login string) (*entity.User, error)
}

type UserByCredentialGetter interface {
	GetByCredential(ctx context.Context, login, passwordHash string) (*entity.User, error)
}

type UserAdder interface {
	Add(ctx context.Context, user *entity.User) error
}

type UserChanger interface {
	Change(ctx context.Context, user *entity.User) error
}

type UserPasswordChanger interface {
	ChangePassword(ctx context.Context, user *entity.User) error
}

type UserRemover interface {
	Remove(ctx context.Context, userID uuid.UUID) error
}

type UserAccrualer interface {
	Accrual(ctx context.Context, user *entity.User, accrual float32) error
}

type UserWithdrawer interface {
	Withdraw(ctx context.Context, user *entity.User, sum float32) error
}

type UserAddWithdrawer interface {
	UserGetter
	UserWithdrawer
}

type UserTransactionalAccrualer interface {
	TransactSupporter
	UserAccrualer
}

type UserRegistrator interface {
	UserByLoginGetter
	UserAdder
}

type UserRepository interface {
	Transactor
	UserGetter
	UserByLoginGetter
	UserByCredentialGetter
	UserAdder
	UserChanger
	UserPasswordChanger
	UserRemover
	UserAccrualer
	UserWithdrawer
}
