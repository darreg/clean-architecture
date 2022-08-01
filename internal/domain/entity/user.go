package entity

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID
	Login        string
	PasswordHash string
	Current      float32
}

func NewUser(login, passwordHash string) *User {
	return &User{
		ID:           uuid.New(),
		Login:        login,
		PasswordHash: passwordHash,
		Current:      0,
	}
}
