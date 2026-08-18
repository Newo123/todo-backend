package middleware

import (
	"net/http"

	"github.com/Newo123/todo-backend/internal/infrastructure/logger"
)

// Logger — middleware, кладущий логгер в контекст запроса.
// Обогащает логгер полями request_id и url, чтобы все последующие
// обработчики автоматически логировали эти поля.
func Logger(log logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)

			l := log.With(
				logger.String("request_id", requestID),
				logger.String("url", r.URL.String()),
			)

			ctx := logger.WithContext(r.Context(), l)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
