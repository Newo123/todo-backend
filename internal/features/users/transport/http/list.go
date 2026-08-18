package http

import (
	"net/http"

	"github.com/Newo123/todo-backend/internal/features/users"
	"github.com/Newo123/todo-backend/internal/infrastructure/logger"
	"github.com/Newo123/todo-backend/internal/transport/http/request"
	"github.com/Newo123/todo-backend/internal/transport/http/response"
)

type ListResponse []UserDTOResponse

func (h *HTTPHandler) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	response := response.NewResponse(log, w)
	request := request.NewRequest(r)

	limit, err := request.QueryInt("limit")
	if err != nil {
		response.Error(err, "failed to get 'limit' query param")
		return
	}
	offset, err := request.QueryInt("offset")
	if err != nil {
		response.Error(err, "failed to get 'offset' query param")
		return
	}

	serviceParams := users.ListParams{
		Limit:  limit,
		Offset: offset,
	}
	serviceResult, err := h.service.List(request.Context(), serviceParams)
	if err != nil {
		response.Error(err, "failed to get users")
		return
	}

	data := ListResponse(usersDTOFromDomains(serviceResult.Users))

	response.OK(data)
}
