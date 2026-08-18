package hasher

import (
	"fmt"

	"github.com/alexedwards/argon2id"
)

type Argon2IDHasher struct {
	params *argon2id.Params
}

func NewArgon2IDHasher(params *argon2id.Params) *Argon2IDHasher {
	if params == nil {
		params = argon2id.DefaultParams
	}

	return &Argon2IDHasher{
		params: params,
	}
}

func (h *Argon2IDHasher) HashPassword(password string) (string, error) {
	if password == "" {
		return "", ErrInvalidPassword
	}

	hash, err := argon2id.CreateHash(password, h.params)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrHashFailed, err)
	}

	return hash, nil
}

func (h *Argon2IDHasher) CheckPassword(hashedPassword, password string) error {
	if hashedPassword == "" || password == "" {
		return ErrInvalidPassword
	}

	match, err := argon2id.ComparePasswordAndHash(password, hashedPassword)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrHashFailed, err)
	}

	if !match {
		return ErrInvalidPassword
	}

	return nil
}
