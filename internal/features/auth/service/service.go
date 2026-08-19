package service

import (
	"github.com/Newo123/todo-backend/internal/features/users"
	"github.com/Newo123/todo-backend/internal/infrastructure/jwt"
)

type Service struct {
	usersService users.Service
	jwtManager   jwt.JWTManager
}

func NewService(
	usersService users.Service,
	jwtManager jwt.JWTManager,
) *Service {
	return &Service{
		usersService: usersService,
		jwtManager:   jwtManager,
	}
}
