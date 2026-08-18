package request

import (
	"fmt"
	"strconv"

	"github.com/Newo123/todo-backend/internal/domain"
	"github.com/google/uuid"
)

// pathValue — универсальный парсер для path-параметров
func (r *Request) pathValue(key string, parser func(string) (any, error)) (any, error) {
	raw := r.r.PathValue(key)
	if raw == "" {
		return nil, fmt.Errorf("path param '%s' is missing: %w", key, domain.ErrInvalidArgument)
	}

	val, err := parser(raw)
	if err != nil {
		return nil, fmt.Errorf("invalid path param '%s'='%s': %w", key, raw, domain.ErrInvalidArgument)
	}
	return val, nil
}

// PathInt извлекает int из path
func (r *Request) PathInt(key string) (int, error) {
	val, err := r.pathValue(key, func(s string) (any, error) {
		return strconv.Atoi(s)
	})
	if err != nil {
		return 0, err
	}
	return val.(int), nil
}

// PathInt извлекает string из path
func (r *Request) PathString(key string) (string, error) {
	val := r.r.PathValue(key)
	if val == "" {
		return "", fmt.Errorf("path param '%s' is empty: %w", key, domain.ErrInvalidArgument)
	}
	return val, nil
}

// PathUUID извлекает UUID из path
// (например, /users/{id})
func (r *Request) PathUUID(key string) (uuid.UUID, error) {
	val, err := r.pathValue(key, func(s string) (any, error) {
		return uuid.Parse(s)
	})
	if err != nil {
		return uuid.Nil, err
	}
	return val.(uuid.UUID), nil
}
