package http

import (
	"net/http"

	"github.com/Newo123/todo-backend/internal/features/users"
	"github.com/Newo123/todo-backend/internal/infrastructure/logger"
	"github.com/Newo123/todo-backend/internal/transport/http/request"
	"github.com/Newo123/todo-backend/internal/transport/http/response"
)

func (h *HTTPHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	response := response.NewResponse(log, w)
	request := request.NewRequest(r)

	userID, err := request.PathUUID("id")
	if err != nil {
		response.Error(err, "failed to get userID")
		return
	}

	serviceParams := users.DeleteParams{ID: userID}
	if err := h.service.Delete(request.Context(), serviceParams); err != nil {
		response.Error(err, "failed to delete user")
		return
	}

	response.NoContent()
}
