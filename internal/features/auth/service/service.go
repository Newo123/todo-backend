package service

import (
	"github.com/Newo123/todo-backend/internal/features/mail"
	"github.com/Newo123/todo-backend/internal/features/users"
	"github.com/Newo123/todo-backend/internal/infrastructure/jwt"
)

type Service struct {
	usersService users.Service
	jwtManager   jwt.JWTManager
	mailService  mail.Service
}

func NewService(
	usersService users.Service,
	jwtManager jwt.JWTManager,
	mailService mail.Service,
) *Service {
	return &Service{
		usersService: usersService,
		jwtManager:   jwtManager,
		mailService:  mailService,
	}
}
