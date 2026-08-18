package request

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Newo123/todo-backend/internal/domain"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

// Response инкапсулирует логику записи HTTP-запросов.
type Request struct {
	r *http.Request
}

func NewRequest(r *http.Request) *Request {
	return &Request{r: r}
}

// Context возвращает контекст запроса
func (r *Request) Context() context.Context {
	return r.r.Context()
}

// Raw возвращает оригинальный http.Request (если очень нужно)
func (r *Request) Raw() *http.Request {
	return r.r
}

// BindJSON читает и валидирует JSON-тело запроса
func (r *Request) BindJSON(dest any) error {
	if err := json.NewDecoder(r.r.Body).Decode(dest); err != nil {
		return fmt.Errorf("decode json: %w", domain.ErrInvalidArgument)
	}
	if err := validate.Struct(dest); err != nil {
		return fmt.Errorf("validation failed: %w", domain.ErrInvalidArgument)
	}
	return nil
}
