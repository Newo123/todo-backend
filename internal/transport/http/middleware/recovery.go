package middleware

import (
	"net/http"

	"github.com/Newo123/todo-backend/internal/infrastructure/logger"
	"github.com/Newo123/todo-backend/internal/transport/http/response"
)

// Panic — middleware для перехвата паник и возврата HTTP 500.
// Без этого middleware паника в обработчике уронила бы всю горутину,
// а стандартная библиотека Go вернула бы пустой ответ клиенту.
//
// Использует defer + recover — стандартный паттерн обработки паник в Go.
func Recovery() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := logger.FromContext(ctx)
			response := response.NewResponse(log, w)

			defer func() {
				if p := recover(); p != nil {
					response.Panic(p, "during handle HTTP request got unexpected panic")
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
