package http

import (
	"net/http"

	"github.com/Newo123/todo-backend/internal/features/auth"
	"github.com/Newo123/todo-backend/internal/transport/http/server"
)

type HTTPHandler struct {
	service auth.Service
}

func NewHTTPHandler(service auth.Service) *HTTPHandler {
	return &HTTPHandler{
		service: service,
	}
}

func (h *HTTPHandler) Routes() []server.Route {
	return []server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/auth/register",
			Handler: h.Register,
		},
		{
			Method:  http.MethodPost,
			Path:    "/auth/login",
			Handler: h.Register,
		},
		{
			Method:  http.MethodPost,
			Path:    "/auth/refresh",
			Handler: h.Register,
		},
	}
}
