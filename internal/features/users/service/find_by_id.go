package service

import (
	"context"

	"github.com/Newo123/todo-backend/internal/features/users"
)

func (s *Service) FindByID(ctx context.Context, params users.FindByIDParams) (users.FindByIDResult, error) {
	user, err := s.repo.FindByID(ctx, params.ID)
	if err != nil {
		return users.FindByIDResult{}, err
	}

	return users.FindByIDResult{User: user}, nil
}
