package adapter

import (
	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/google/uuid"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u UserRepository) Find(id uuid.UUID) *entity.User {
	// TODO implement me
	panic("implement me")
}

func (u UserRepository) FindByLogin(login string) *entity.User {
	// TODO implement me
	panic("implement me")
}
