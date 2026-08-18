package request

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Newo123/todo-backend/internal/domain"
	"github.com/google/uuid"
)

// queryParam — универсальный парсер для query-параметров
// (возвращает nil, если параметр отсутствует)
func (r *Request) queryParam(key string, parser func(string) (any, error)) (any, error) {
	raw := r.r.URL.Query().Get(key)
	if raw == "" {
		return nil, nil
	}

	val, err := parser(raw)
	if err != nil {
		return nil, fmt.Errorf("invalid query param '%s'='%s': %w", key, raw, domain.ErrInvalidArgument)
	}
	return val, nil
}

// QueryBool извлекает bool из query-параметра
// (true/false)
func (r *Request) QueryBool(key string) (*bool, error) {
	val, err := r.queryParam(key, func(s string) (any, error) {
		return strconv.ParseBool(s)
	})
	if err != nil {
		return nil, err
	}
	if val == "" {
		return nil, nil
	}
	parsed := val.(bool)
	return &parsed, nil
}

// QueryString извлекает строку из query-параметра
// (возвращает nil, если параметр отсутствует)
func (r *Request) QueryString(key string) (*string, error) {
	val := r.r.URL.Query().Get(key)
	if val == "" {
		return nil, nil
	}
	return &val, nil
}

// QueryDate извлекает дату из query-параметра формата YYYY-MM-DD
func (r *Request) QueryDate(key string) (*time.Time, error) {
	val, err := r.queryParam(key, func(s string) (any, error) {
		return time.Parse(time.DateOnly, s)
	})
	if err != nil {
		return nil, err
	}
	if val == nil {
		return nil, nil
	}
	parsed := val.(time.Time)
	return &parsed, nil
}

// QueryInt извлекает int из query-параметра
// (возвращает nil, если параметр отсутствует)
func (r *Request) QueryInt(key string) (*int, error) {
	val, err := r.queryParam(key, func(s string) (any, error) {
		return strconv.Atoi(s)
	})
	if err != nil {
		return nil, err
	}
	if val == nil {
		return nil, nil
	}
	parsed := val.(int)
	return &parsed, nil
}

// QueryUUID извлекает UUID из query-параметра
// (возвращает nil, если параметр отсутствует)
func (r *Request) QueryUUID(key string) (*uuid.UUID, error) {
	val, err := r.queryParam(key, func(s string) (any, error) {
		return uuid.Parse(s)
	})
	if err != nil {
		return nil, err
	}
	if val == nil {
		return nil, nil
	}
	parsed := val.(uuid.UUID)
	return &parsed, nil
}
