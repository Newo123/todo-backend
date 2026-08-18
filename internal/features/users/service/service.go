package service

import (
	"github.com/Newo123/todo-backend/internal/features/users"
	"github.com/Newo123/todo-backend/internal/infrastructure/hasher"
)

type Service struct {
	repo   users.Repository
	hasher hasher.Hasher
}

func NewService(repo users.Repository, hasher hasher.Hasher) *Service {
	return &Service{repo: repo, hasher: hasher}
}
