package users

import (
	"context"

	"github.com/Newo123/todo-backend/internal/domain"
	"github.com/google/uuid"
)

type Repository interface {
	Create(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
	FindByID(
		ctx context.Context,
		id uuid.UUID,
	) (domain.User, error)
	List(
		ctx context.Context,
		limit *int,
		offset *int,
	) ([]domain.User, error)
	Delete(
		ctx context.Context,
		id uuid.UUID,
	) error
}
