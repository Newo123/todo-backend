package service

import (
	"context"

	"github.com/Newo123/todo-backend/internal/features/auth"
	"github.com/Newo123/todo-backend/internal/features/users"
)

func (s *Service) Register(ctx context.Context, params auth.RegisterParams) error {

	/*
		1. Создаем пользователя
		2. Публикуем событие о создании пользователя
		3. Отдаем nil ответ
	*/
	usersServiceParams := users.CreateParams{
		Email:    params.Email,
		Password: params.Password,
		FullName: params.FullName,
	}
	user, err := s.usersService.Create(ctx, usersServiceParams)
	if err != nil {
		return err
	}

}
