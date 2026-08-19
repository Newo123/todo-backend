package mail

import (
	"context"

	"github.com/Newo123/todo-backend/internal/domain"
)

type Service interface {
	SendVerificationEmail(ctx context.Context, user domain.User) error
}
