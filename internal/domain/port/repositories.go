package port

import (
	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/google/uuid"
)

type UserRepository interface {
	Find(userID uuid.UUID) (*entity.User, error)
	FindByLogin(login string) (*entity.User, error)
	FindByCredential(login, password string) (*entity.User, error)
	Create(login, password string) (*entity.User, error)
	Update(login, password string) (*entity.User, error)
	Delete(userID uuid.UUID) error
}

type OrderRepository interface {
	Find(number string) (*entity.Order, error)
	FindByUser(user entity.User) (*entity.Order, error)
}
