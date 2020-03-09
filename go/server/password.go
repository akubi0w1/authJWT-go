package server

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(pw string) (string, error) {
	if len(pw) > 70 {
		return "", errors.New("password is too long.")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func PasswordVerify(hash, pw string) error {
	// 認証に失敗した場合は error
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
}
