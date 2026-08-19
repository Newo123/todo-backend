package auth

import (
	"context"

	"github.com/Newo123/todo-backend/internal/domain"
)

type Service interface {
	Register(
		ctx context.Context,
		params RegisterParams,
	) error
	Login(
		ctx context.Context,
		params LoginParams,
	) (LoginResult, error)
	RefreshToken(
		ctx context.Context,
	) error
}

type RegisterParams struct {
	Email    string
	Password string
	FullName *string
}

// type RegisterResult struct {
// Message string
// AccessToken     string
// RefreshToken    string
// RefreshTokenTTL time.Duration
// }

// Login Struct

type LoginParams struct {
	Email    string
	Password string
}
type LoginResult struct {
	User        domain.User
	AccessToken string
}
