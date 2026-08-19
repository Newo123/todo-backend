package jwt

import "errors"

var (
	ErrInvalidUserID        = errors.New("invalid user id")
	ErrInvalidToken         = errors.New("invalid token")
	ErrInvalidTokenType     = errors.New("invalid token type")
	ErrInvalidSigningMethod = errors.New("invalid signing method")
	ErrTokenGeneration      = errors.New("token generation failed")
)
