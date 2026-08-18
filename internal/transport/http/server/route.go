package server

import (
	"net/http"

	"github.com/Newo123/todo-backend/internal/transport/http/middleware"
)

// Route описывает один HTTP-маршрут: метод, путь, обработчик и middleware.
// Middleware в Route применяются только к этому конкретному маршруту,
// в отличие от middleware сервера (применяются ко всем маршрутам).
type Route struct {
	Method     string
	Path       string
	Handler    http.HandlerFunc
	Middleware []middleware.Middleware
}

// WithMiddleware применяет middleware маршрута к обработчику и возвращает готовый http.Handler.
func (r *Route) WithMiddleware() http.Handler {
	return middleware.ChainMiddleware(r.Handler, r.Middleware...)
}
