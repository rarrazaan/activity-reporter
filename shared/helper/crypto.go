package helper

import (
	"golang.org/x/crypto/bcrypt"
)

type AppCrypto interface {
	HashAndSalt(pwd []byte) (string, error)
	ComparePasswords(password, hashedPassword string) error
}

type appCrypto struct {
}

func NewAppCrypto() *appCrypto {
	return &appCrypto{}
}

// HashAndSalt Hashes a given string
func (d appCrypto) HashAndSalt(pwd []byte) (string, error) {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (d appCrypto) ComparePasswords(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
