package hasher

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct {
	cost int
}

func NewBcryptHasher(cost int) *BcryptHasher {
	if cost == 0 {
		cost = bcrypt.DefaultCost
	}
	return &BcryptHasher{cost: cost}
}

func (h *BcryptHasher) HashPassword(password string) (string, error) {
	if password == "" {
		return "", ErrInvalidPassword
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrHashFailed, err)
	}
	return string(hash), nil
}

func (h *BcryptHasher) CheckPassword(hashedPassword, password string) error {
	if hashedPassword == "" || password == "" {
		return ErrInvalidPassword
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return ErrInvalidPassword
		}
		return fmt.Errorf("%w: %w", ErrHashFailed, err)
	}

	return nil
}
