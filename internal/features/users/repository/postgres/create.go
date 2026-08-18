package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Newo123/todo-backend/internal/domain"
	"github.com/Newo123/todo-backend/internal/infrastructure/postgres"
)

func (r *Repository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	row := r.pool.QueryRow(
		ctx,
		insertQuery,
		user.ID,
		user.Version,
		user.Email,
		user.EmailVerified,
		user.PasswordHash,
		user.FullName,
		user.CreatedAt,
		user.UpdatedAt,
	)

	var userModel UserModel
	if err := userModel.Scan(row); err != nil {
		if errors.Is(err, postgres.ErrViolatesUniqueKey) {
			return domain.User{}, domain.ErrAlreadyExists
		}

		return domain.User{}, fmt.Errorf("scan error")
	}

	userDomain := modelToDomain(userModel)

	return userDomain, nil
}
