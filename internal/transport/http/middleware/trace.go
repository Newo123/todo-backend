package middleware

import (
	"net/http"
	"time"

	"github.com/Newo123/todo-backend/internal/infrastructure/logger"
	"github.com/Newo123/todo-backend/internal/transport/http/response"
)

// Trace — middleware для логирования входящих запросов и времени их обработки.
// Использует ResponseWriter-обёртку, чтобы перехватить статус-код ответа,
// который иначе недоступен после вызова WriteHeader.
func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := logger.FromContext(ctx)

			rw := response.NewResponseWriter(w)

			before := time.Now()
			log.Debug(
				">>> incoming HTTP request",
				logger.String("http_method", r.Method),
				logger.Time("time", before.UTC()),
			)

			next.ServeHTTP(rw, r)

			log.Debug(
				"<<< done HTTP request",
				logger.Int("status_code", rw.GetStatusCode()),
				logger.Duration("latency", time.Since(before)),
			)
		})
	}
}
