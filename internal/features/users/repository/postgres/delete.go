package postgres

import (
	"context"
	"fmt"

	"github.com/Newo123/todo-backend/internal/domain"
	"github.com/google/uuid"
)

func (r *Repository) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	cmdTag, err := r.pool.Exec(ctx, deleteQuery, id)
	if err != nil {
		return fmt.Errorf("exec query")
	}
	if cmdTag.RowsAffected() == 0 {
		return domain.ErrNotFound
	}

	return nil
}
