package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/alrund/yp-1-project/internal/application/usecase"
	"github.com/alrund/yp-1-project/internal/domain/entity"
	"github.com/alrund/yp-1-project/internal/infrastructure/helper"
	"github.com/google/uuid"
)

type UserRepository struct {
	helper.Transaction
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
	return u.InTransaction(ctx, u.db, tFunc)
}

func (u UserRepository) Get(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	var user entity.User

	err := u.QueryRowContext(ctx, u.db,
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

func (u UserRepository) GetByLogin(ctx context.Context, login string) (*entity.User, error) {
	var user entity.User

	err := u.QueryRowContext(ctx, u.db,
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

func (u UserRepository) GetByCredential(ctx context.Context, login, passwordHash string) (*entity.User, error) {
	var user entity.User

	err := u.QueryRowContext(ctx, u.db,
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

func (u UserRepository) Withdraw(ctx context.Context, user *entity.User, sum float32) error {
	_, err := u.ExecContext(ctx, u.db,
		"UPDATE users SET current=current-$2 WHERE id=$1", user.ID, sum,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) Add(ctx context.Context, user *entity.User) error {
	_, err := u.ExecContext(ctx, u.db,
		"INSERT INTO users(ID, login, password) VALUES($1, $2, $3)",
		user.ID, user.Login, user.PasswordHash,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) Change(ctx context.Context, user *entity.User) error {
	_, err := u.ExecContext(ctx, u.db,
		"UPDATE users SET login=$2, current=$3 WHERE id=$1", user.ID, user.Login, user.Current,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) ChangePassword(ctx context.Context, user *entity.User) error {
	_, err := u.ExecContext(ctx, u.db, "UPDATE users SET password=$2 WHERE id=$1", user.ID, user.PasswordHash)
	if err != nil {
		return err
	}

	return nil
}

func (u UserRepository) Remove(ctx context.Context, userID uuid.UUID) error {
	_, err := u.ExecContext(ctx, u.db, "DELETE FROM users WHERE id=$1", userID)
	return err
}
