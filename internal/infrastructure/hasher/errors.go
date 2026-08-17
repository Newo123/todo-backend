package hasher

import "errors"

var (
	ErrInvalidPassword = errors.New("invalid password")
	ErrHashFailed      = errors.New("failed to hash password")
	ErrInvalidHash     = errors.New("invalid hash format")
)
