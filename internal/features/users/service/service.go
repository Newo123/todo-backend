package service

import (
	"github.com/Newo123/todo-backend/internal/features/mail"
	"github.com/Newo123/todo-backend/internal/features/users"
	"github.com/Newo123/todo-backend/internal/infrastructure/hasher"
)

type Service struct {
	repo        users.Repository
	hasher      hasher.Hasher
	mailService mail.Service
}

func NewService(
	repo users.Repository,
	hasher hasher.Hasher,
	mailService mail.Service,
) *Service {
	return &Service{
		repo:        repo,
		hasher:      hasher,
		mailService: mailService,
	}
}
