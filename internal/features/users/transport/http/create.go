package http

import (
	"net/http"

	"github.com/Newo123/todo-backend/internal/features/users"
	"github.com/Newo123/todo-backend/internal/infrastructure/logger"
	"github.com/Newo123/todo-backend/internal/transport/http/request"
	"github.com/Newo123/todo-backend/internal/transport/http/response"
)

type CreateRequest struct {
	Email         string  `json:"email" validate:"required,email"`
	EmailVerified bool    `json:"email_verified"`
	Password      string  `json:"password" validate:"required,min=6,max=100"`
	FullName      *string `json:"full_name" validate:"omitempty,min=3,max=100"`
}

type CreateResponse UserDTOResponse

func (h *HTTPHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	response := response.NewResponse(log, w)
	request := request.NewRequest(r)

	var body CreateRequest
	if err := request.BindJSON(&body); err != nil {
		response.Error(err, "failed to decode and validate HTTP request")
		return
	}

	serviceParams := users.CreateParams{
		Email:         body.Email,
		Password:      body.Password,
		FullName:      body.FullName,
		EmailVerified: body.EmailVerified,
	}
	serviceResult, err := h.service.Create(request.Context(), serviceParams)
	if err != nil {
		response.Error(err, "failed to create user")
		return
	}

	data := CreateResponse(userDTOFromDomain(serviceResult.User))

	response.OK(data)
}
