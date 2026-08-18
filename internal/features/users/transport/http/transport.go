package http

import (
	"net/http"

	"github.com/Newo123/todo-backend/internal/features/users"
	"github.com/Newo123/todo-backend/internal/transport/http/server"
)

type HTTPHandler struct {
	service users.Service
}

func NewHTTPTransport(service users.Service) *HTTPHandler {
	return &HTTPHandler{service: service}
}

func (h *HTTPHandler) Routes() []server.Route {
	return []server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.Create,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: h.FindByID,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: h.List,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{id}",
			Handler: h.Delete,
		},
	}
}
