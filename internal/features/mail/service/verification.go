package service

import (
	"bytes"
	"context"
	"fmt"
	"html/template"

	"github.com/Newo123/todo-backend/internal/domain"
	"github.com/Newo123/todo-backend/internal/features/mail"
	infrastructureMail "github.com/Newo123/todo-backend/internal/infrastructure/mail"
)

func (s *Service) SendVerificationEmail(
	ctx context.Context,
	user domain.User,
) error {
	tmpl, err := template.New("verification").Parse(mail.VerificationTemplate)
	if err != nil {
		return fmt.Errorf("parse verification email template: %w", err)
	}

	var body bytes.Buffer

	if err := tmpl.Execute(&body, user); err != nil {
		return fmt.Errorf("execute verification email template: %w", err)
	}

	message := infrastructureMail.Message{
		To:      user.Email,
		Subject: "Подтверждение!",
		Body:    body.String(),
	}
	if err := s.mailer.Send(ctx, message); err != nil {
		return fmt.Errorf("send verification email: %w", err)
	}

	return nil
}
