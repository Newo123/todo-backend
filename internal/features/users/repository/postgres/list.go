package postgres

import (
	"context"
	"fmt"

	"github.com/Newo123/todo-backend/internal/domain"
)

func (r *Repository) List(ctx context.Context, limit *int, offset *int) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	rows, err := r.pool.Query(
		ctx,
		listQuery,
		limit,
		offset,
	)
	if err != nil {
		return nil, fmt.Errorf("select users")
	}
	defer rows.Close()

	var userModels []UserModel
	for rows.Next() {
		var userModel UserModel
		if err := userModel.Scan(rows); err != nil {
			return nil, fmt.Errorf("scan users")
		}

		userModels = append(userModels, userModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows")
	}

	userDomains := modelsToDomains(userModels)

	return userDomains, nil
}
