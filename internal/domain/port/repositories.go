package port

import (
	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/google/uuid"
)

type UserRepository interface {
	Find(id uuid.UUID) *entity.User
	FindByLogin(login string) *entity.User
}

type OrderRepository interface {
	Find(number string) *entity.Order
	FindByUser(user entity.User) *entity.Order
}
