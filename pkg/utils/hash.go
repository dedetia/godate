package utils

import (
	"errors"
	"github.com/matthewhartstonge/argon2"
)

const (
	AlgArgon = "argon"
)

func Hash(alg, password string) ([]byte, error) {
	switch alg {
	case AlgArgon:
		argon := argon2.DefaultConfig()
		h, err := argon.HashEncoded([]byte(password))
		if err != nil {
			return nil, err
		}
		return h, nil
	default:
		return nil, errors.New("unsupported algorithm")
	}
}

func Verify(alg, password string, hashedPassword []byte) error {
	switch alg {
	case AlgArgon:
		passwordBytes := []byte(password)
		ok, err := argon2.VerifyEncoded(passwordBytes, hashedPassword)
		if err != nil || !ok {
			return errors.New("invalid password")
		}
		return nil
	default:
		return errors.New("unsupported algorithm")
	}
}
