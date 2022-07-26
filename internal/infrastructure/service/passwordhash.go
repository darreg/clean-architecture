package service

import (
	"crypto/sha256"
	"fmt"
)

type PasswordHash struct{}

func NewPasswordHasher() *PasswordHash {
	return &PasswordHash{}
}

func (h *PasswordHash) Hash(password string) string {
	pwd := sha256.New()
	pwd.Write([]byte(password))

	return fmt.Sprintf("%x", pwd.Sum(nil))
}
