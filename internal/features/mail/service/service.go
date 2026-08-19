package service

import "github.com/Newo123/todo-backend/internal/infrastructure/mail"

type Service struct {
	mailer mail.Mailer
}

func NewService(
	mailer mail.Mailer,
) *Service {
	return &Service{mailer: mailer}
}
