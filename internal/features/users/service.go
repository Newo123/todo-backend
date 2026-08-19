package users

import (
	"context"

	"github.com/Newo123/todo-backend/internal/domain"
	"github.com/google/uuid"
)

type Service interface {
	Create(
		ctx context.Context,
		params CreateParams,
	) (CreateResult, error)
	FindByID(
		ctx context.Context,
		params FindByIDParams,
	) (FindByIDResult, error)
	List(
		ctx context.Context,
		params ListParams,
	) (ListResult, error)
	Delete(
		ctx context.Context,
		params DeleteParams,
	) error
}

// Create
type CreateParams struct {
	Email         string
	Password      string
	FullName      *string
	EmailVerified bool
}
type CreateResult struct {
	User domain.User
}

// Find By ID
type FindByIDParams struct {
	ID uuid.UUID
}
type FindByIDResult struct {
	User domain.User
}

// List
type ListParams struct {
	Limit  *int
	Offset *int
}
type ListResult struct {
	Users []domain.User
}

// Delete
type DeleteParams struct {
	ID uuid.UUID
}
