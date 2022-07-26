package adapter

import (
	"database/sql"
	"errors"

	"github.com/alrund/yp-1-project/internal/application/usecase"
	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/alrund/yp-1-project/internal/domain/port"
	"github.com/google/uuid"
)

type UserRepository struct {
	db     *sql.DB
	hasher port.PasswordHasher
}

func NewUserRepository(db *sql.DB, hasher port.PasswordHasher) *UserRepository {
	return &UserRepository{db: db, hasher: hasher}
}

func (u UserRepository) Find(id uuid.UUID) (*entity.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserRepository) FindByLogin(login string) (*entity.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u UserRepository) FindByCredential(login, password string) (*entity.User, error) {
	var (
		uID           uuid.UUID
		ulogin        string
		upasswordHash string
	)

	err := u.db.QueryRow(
		"SELECT * FROM users WHERE login = $1 AND password = $2", login, u.hasher.Hash(password),
	).Scan(&uID, &ulogin, &upasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrUserNotFound
		}
		return nil, err
	}

	return &entity.User{ID: uID, Login: ulogin, PasswordHash: upasswordHash}, nil
}

func (u UserRepository) Create(login, password string) (*entity.User, error) {
	panic("implement me")
}

func (u UserRepository) Update(login, password string) (*entity.User, error) {
	panic("implement me")
}

func (u UserRepository) Delete(userID uuid.UUID) error {
	panic("implement me")
}
