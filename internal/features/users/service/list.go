package service

import (
	"context"
	"fmt"

	"github.com/Newo123/todo-backend/internal/domain"
	"github.com/Newo123/todo-backend/internal/features/users"
)

func (s *Service) List(ctx context.Context, params users.ListParams) (users.ListResult, error) {
	if params.Limit != nil && *params.Limit < 0 {
		return users.ListResult{}, fmt.Errorf(
			"limit must be non-negative: %w",
			domain.ErrInvalidArgument,
		)
	}

	if params.Offset != nil && *params.Offset < 0 {
		return users.ListResult{}, fmt.Errorf(
			"offset must be non-negative: %w",
			domain.ErrInvalidArgument,
		)
	}

	userList, err := s.repo.List(ctx, params.Limit, params.Offset)
	if err != nil {
		return users.ListResult{}, err
	}

	return users.ListResult{Users: userList}, nil
}
