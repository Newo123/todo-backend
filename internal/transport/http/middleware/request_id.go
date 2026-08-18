package middleware

import (
	"net/http"

	"github.com/google/uuid"
)

var requestIDHeader = "X-Request-ID"

// RequestID — middleware, обеспечивающий каждый запрос уникальным идентификатором.
// Если клиент передаёт X-Request-ID — используем его (полезно для распределённой трассировки).
func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)
			if requestID == "" {
				requestID = uuid.NewString()
			}

			r.Header.Set(requestIDHeader, requestID)
			w.Header().Set(requestIDHeader, requestID)

			next.ServeHTTP(w, r)
		})
	}
}
