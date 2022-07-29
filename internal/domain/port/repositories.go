package port

import (
	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/google/uuid"
)

type UserRepository interface {
	Get(userID uuid.UUID) (*entity.User, error)
	GetByLogin(login string) (*entity.User, error)
	GetByCredential(login, passwordHash string) (*entity.User, error)
	Add(user *entity.User) error
	Change(user *entity.User) error
	ChangePassword(user *entity.User) error
	Remove(userID uuid.UUID) error
}

type OrderRepository interface {
	Get(number string) (*entity.Order, error)
	GetAllByUser(user *entity.User) ([]*entity.Order, error)
	Add(order *entity.Order) error
	Change(order *entity.Order) error
	Remove(number string) error
}

type WithdrawRepository interface {
	Get(withdrawID uuid.UUID) (*entity.Withdraw, error)
	GetAllByUser(user *entity.User) ([]*entity.Withdraw, error)
	Add(order *entity.Withdraw) error
	Change(order *entity.Withdraw) error
	Remove(withdrawID uuid.UUID) error
}
