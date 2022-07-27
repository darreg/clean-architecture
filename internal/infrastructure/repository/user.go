package repository

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

func (u UserRepository) Get(userID uuid.UUID) (*entity.User, error) {
	var user entity.User

	err := u.db.QueryRow(
		"SELECT * FROM users WHERE id = $1", userID,
	).Scan(&user.ID, &user.Login, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (u UserRepository) GetByLogin(login string) (*entity.User, error) {
	var user entity.User

	err := u.db.QueryRow(
		"SELECT * FROM users WHERE login = $1", login,
	).Scan(&user.ID, &user.Login, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (u UserRepository) GetByCredential(login, password string) (*entity.User, error) {
	var user entity.User

	err := u.db.QueryRow(
		"SELECT * FROM users WHERE login = $1 AND password=$2", login, u.hasher.Hash(password),
	).Scan(&user.ID, &user.Login, &user.PasswordHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (u UserRepository) Add(userID uuid.UUID, login, password string) (*entity.User, error) {
	passwordHash := u.hasher.Hash(password)

	_, err := u.db.Exec("INSERT INTO users(ID, login, password) VALUES($1, $2, $3)", userID, login, passwordHash)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:           userID,
		Login:        login,
		PasswordHash: passwordHash,
	}, nil
}

func (u UserRepository) Change(userID uuid.UUID, login, password string) (*entity.User, error) {
	passwordHash := u.hasher.Hash(password)

	_, err := u.db.Exec("UPDATE users SET login=$1, password=$2 WHERE id =$3", login, passwordHash, userID)
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:           userID,
		Login:        login,
		PasswordHash: passwordHash,
	}, nil
}

func (u UserRepository) Remove(userID uuid.UUID) error {
	_, err := u.db.Exec("DELETE FROM users WHERE id=$1", userID)
	return err
}
