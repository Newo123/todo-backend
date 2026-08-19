package service

import (
	"context"

	"github.com/Newo123/todo-backend/internal/domain"
	"github.com/Newo123/todo-backend/internal/features/users"
)

func (s *Service) Create(ctx context.Context, params users.CreateParams) (users.CreateResult, error) {
	passwordHash, err := s.hasher.HashPassword(params.Password)
	if err != nil {
		return users.CreateResult{}, err
	}

	user, err := domain.CreateUser(
		params.Email,
		passwordHash,
		params.FullName,
		params.EmailVerified,
	)
	if err != nil {
		return users.CreateResult{}, err
	}

	user, err = s.repo.Create(ctx, user)
	if err != nil {
		return users.CreateResult{}, err
	}

	return users.CreateResult{User: user}, nil
}
