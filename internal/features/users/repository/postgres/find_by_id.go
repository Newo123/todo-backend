package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/Newo123/todo-backend/internal/domain"
	"github.com/Newo123/todo-backend/internal/infrastructure/postgres"
	"github.com/google/uuid"
)

func (r *Repository) FindByID(ctx context.Context, id uuid.UUID) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	row := r.pool.QueryRow(ctx, findByIDQuery, id)

	var userModel UserModel
	if err := userModel.Scan(row); err != nil {
		if errors.Is(err, postgres.ErrNoRows) {
			return domain.User{}, domain.ErrNotFound
		}

		return domain.User{}, fmt.Errorf("scan error")
	}

	userDomain := modelToDomain(userModel)

	return userDomain, nil
}
