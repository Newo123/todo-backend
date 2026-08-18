package service

import (
	"context"

	"github.com/Newo123/todo-backend/internal/features/users"
)

func (s *Service) Delete(ctx context.Context, params users.DeleteParams) error {
	if err := s.repo.Delete(ctx, params.ID); err != nil {
		return err
	}

	return nil
}
