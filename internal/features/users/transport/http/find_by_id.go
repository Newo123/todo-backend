package http

import (
	"net/http"

	"github.com/Newo123/todo-backend/internal/features/users"
	"github.com/Newo123/todo-backend/internal/infrastructure/logger"
	"github.com/Newo123/todo-backend/internal/transport/http/request"
	"github.com/Newo123/todo-backend/internal/transport/http/response"
)

type FindByIDResponse UserDTOResponse

func (h *HTTPHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	response := response.NewResponse(log, w)
	request := request.NewRequest(r)

	userID, err := request.PathUUID("id")
	if err != nil {
		response.Error(err, "failed to get userID path value")
		return
	}

	serviceParams := users.FindByIDParams{ID: userID}
	serviceResult, err := h.service.FindByID(request.Context(), serviceParams)
	if err != nil {
		response.Error(err, "failed to get user")
		return
	}

	data := FindByIDResponse(userDTOFromDomain(serviceResult.User))

	response.OK(data)
}
