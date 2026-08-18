package postgres

import "github.com/Newo123/todo-backend/internal/infrastructure/postgres"

type Repository struct {
	pool postgres.Pool
}

func NewRepository(pool postgres.Pool) *Repository {
	return &Repository{pool: pool}
}
