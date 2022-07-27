package port

import (
	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/google/uuid"
)

type UserRepository interface {
	Get(userID uuid.UUID) (*entity.User, error)
	GetByLogin(login string) (*entity.User, error)
	GetByCredential(login, password string) (*entity.User, error)
	Add(userID uuid.UUID, login, password string) (*entity.User, error)
	Change(userID uuid.UUID, login, password string) (*entity.User, error)
	Remove(userID uuid.UUID) error
}

type OrderRepository interface {
	Get(number string) (*entity.Order, error)
	GetByUser(user entity.User) (*entity.Order, error)
}
