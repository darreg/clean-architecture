package repository

import (
	"database/sql"
	"errors"

	"github.com/alrund/yp-1-project/internal/application/usecase"
	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) Get(userID uuid.UUID) (*entity.User, error) {
	var user entity.User

	err := u.db.QueryRow(
		"SELECT id, login, password, current FROM users WHERE id = $1", userID,
	).Scan(&user.ID, &user.Login, &user.PasswordHash, &user.Current)
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
		"SELECT id, login, password, current FROM users WHERE login = $1", login,
	).Scan(&user.ID, &user.Login, &user.PasswordHash, &user.Current)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (u UserRepository) GetByCredential(login, passwordHash string) (*entity.User, error) {
	var user entity.User

	err := u.db.QueryRow(
		"SELECT id, login, password, current FROM users WHERE login = $1 AND password=$2", login, passwordHash,
	).Scan(&user.ID, &user.Login, &user.PasswordHash, &user.Current)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, usecase.ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (u UserRepository) Add(user *entity.User) error {
	_, err := u.db.Exec(
		"INSERT INTO users(ID, login, password) VALUES($1, $2, $3)",
		user.ID, user.Login, user.PasswordHash,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) Change(user *entity.User) error {
	_, err := u.db.Exec("UPDATE users SET login=$2, current=$3 WHERE id=$1", user.ID, user.Login, user.Current)
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) ChangePassword(user *entity.User) error {
	_, err := u.db.Exec("UPDATE users SET password=$2 WHERE id=$1", user.ID, user.PasswordHash)
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) Remove(userID uuid.UUID) error {
	_, err := u.db.Exec("DELETE FROM users WHERE id=$1", userID)
	return err
}
