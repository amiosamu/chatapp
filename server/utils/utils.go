package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashedPassword(pass string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func CheckPassword(pass, hashedPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(pass), []byte(hashedPass))
}
