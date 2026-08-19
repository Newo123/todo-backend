package smtp

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"

	"github.com/Newo123/todo-backend/internal/infrastructure/mail"
)

// Mailer — реализация mail.Mailer через SMTP.
// Содержит конфигурацию, необходимую для подключения
// и аутентификации на SMTP-сервере.
type Mailer struct {
	config Config
}

// NewMailer создаёт SMTP mailer с переданной конфигурацией.
func NewMailer(config Config) *Mailer {
	return &Mailer{
		config: config,
	}
}

// Send отправляет одно письмо через SMTP-сервер.
//
// Если message.From не указан, используется адрес отправителя
// из SMTP-конфигурации.
func (m *Mailer) Send(
	ctx context.Context,
	message mail.Message,
) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	if message.From == "" {
		message.From = m.config.From
	}

	client, err := m.dial(ctx)
	if err != nil {
		return fmt.Errorf("dial SMTP server: %w", err)
	}
	defer client.Close()

	auth := smtp.PlainAuth(
		"",
		m.config.User,
		m.config.Password,
		m.config.Host,
	)

	if err := client.Auth(auth); err != nil {
		return fmt.Errorf("SMTP auth: %w", err)
	}

	if err := client.Mail(message.From); err != nil {
		return fmt.Errorf("SMTP MAIL FROM: %w", err)
	}

	if err := client.Rcpt(message.To); err != nil {
		return fmt.Errorf("SMTP RCPT TO: %w", err)
	}

	writer, err := client.Data()
	if err != nil {
		return fmt.Errorf("SMTP DATA: %w", err)
	}

	if _, err := writer.Write(buildMessage(message)); err != nil {
		_ = writer.Close()

		return fmt.Errorf("write SMTP message: %w", err)
	}

	if err := writer.Close(); err != nil {
		return fmt.Errorf("close SMTP message: %w", err)
	}

	if err := client.Quit(); err != nil {
		return fmt.Errorf("quit SMTP connection: %w", err)
	}

	return nil
}

// dial устанавливает TCP-соединение с SMTP-сервером
// и создаёт поверх него SMTP-клиент.
//
// Использование DialContext позволяет отменить подключение
// при отмене переданного context.Context.
func (m *Mailer) dial(ctx context.Context) (*smtp.Client, error) {
	dialer := net.Dialer{}

	conn, err := dialer.DialContext(
		ctx,
		"tcp",
		m.config.Addr(),
	)
	if err != nil {
		return nil, err
	}

	client, err := smtp.NewClient(conn, m.config.Host)
	if err != nil {
		_ = conn.Close()

		return nil, err
	}

	if ok, _ := client.Extension("STARTTLS"); !ok {
		_ = client.Close()

		return nil, fmt.Errorf("SMTP server does not support STARTTLS")
	}

	if err := client.StartTLS(&tls.Config{
		ServerName: m.config.Host,
	}); err != nil {
		_ = client.Close()

		return nil, fmt.Errorf("SMTP STARTTLS: %w", err)
	}

	return client, nil
}

// buildMessage формирует содержимое SMTP-сообщения.
//
// SMTP требует разделитель строк CRLF (\r\n).
// Content-Type указываем как text/html, поскольку тело письма
// формируется из HTML-шаблонов.
func buildMessage(message mail.Message) []byte {
	return []byte(
		"From: " + message.From + "\r\n" +
			"To: " + message.To + "\r\n" +
			"Subject: " + message.Subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=UTF-8\r\n" +
			"\r\n" +
			message.Body +
			"\r\n",
	)
}
