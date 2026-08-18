package middleware

import "net/http"

// Middleware — тип, описывающий функцию, которая принимает http.Handler
// и возвращает http.Handler (оборачивает его).
type Middleware func(next http.Handler) http.Handler

// Реализация: обходим массив с конца и оборачиваем обработчик снаружи внутрь,
// чтобы первый middleware оказался самым внешним слоем.
func ChainMiddleware(h http.Handler, m ...Middleware) http.Handler {
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}

	return h
}
