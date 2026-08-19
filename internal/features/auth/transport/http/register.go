package http

import (
	"net/http"

	"github.com/Newo123/todo-backend/internal/features/auth"
	"github.com/Newo123/todo-backend/internal/infrastructure/logger"
	"github.com/Newo123/todo-backend/internal/transport/http/request"
	"github.com/Newo123/todo-backend/internal/transport/http/response"
)

type RegisterRequest struct {
	Email    string  `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required,min=6,max=100"`
	FullName *string `json:"full_name" validate:"omitempty,min=3,max=100"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

func (h *HTTPHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := logger.FromContext(ctx)
	response := response.NewResponse(log, w)
	request := request.NewRequest(r)

	var body RegisterRequest
	if err := request.BindJSON(&body); err != nil {
		response.Error(err, "failed to decode and validate HTTP request")
		return
	}

	serviceParams := auth.RegisterParams{
		Email:    body.Email,
		Password: body.Password,
		FullName: body.FullName,
	}
	if err := h.service.Register(request.Context(), serviceParams); err != nil {
		response.Error(err, "failed register user")
		return
	}

	response.Created(RegisterResponse{Message: "На вашу почту отправленно письмо для подтверждения"})
}
