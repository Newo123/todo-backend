package request

import (
	"fmt"

	"github.com/Newo123/todo-backend/internal/domain"
)

// Header извлекает значение заголовка (обязательный)
func (r *Request) Header(key string) (string, error) {
	val := r.r.Header.Get(key)
	if val == "" {
		return "", fmt.Errorf("header '%s' is required: %w", key, domain.ErrInvalidArgument)
	}
	return val, nil
}

// BearerToken извлекает токен из Authorization: Bearer <token>
func (r *Request) BearerToken() (string, error) {
	auth, err := r.Header("Authorization")
	if err != nil {
		return "", err
	}

	const prefix = "Bearer "
	if len(auth) < len(prefix) || auth[:len(prefix)] != prefix {
		return "", fmt.Errorf("invalid Authorization header format: %w", domain.ErrInvalidArgument)
	}

	return auth[len(prefix):], nil
}

// ClientIP возвращает реальный IP клиента (учитывая прокси)
func (r *Request) ClientIP() string {
	if ip := r.r.Header.Get("X-Forwarded-For"); ip != "" {
		return ip
	}
	if ip := r.r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}
	return r.r.RemoteAddr
}
