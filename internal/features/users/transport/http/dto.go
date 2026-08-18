package http

import (
	"time"

	"github.com/Newo123/todo-backend/internal/domain"
	"github.com/google/uuid"
)

// UserDTOResponse — DTO для представления пользователя в API-ответе.
type UserDTOResponse struct {
	ID            uuid.UUID `json:"id"`
	Version       int64     `json:"version"`
	Email         string    `json:"email"`
	EmailVerified bool      `json:"email_verified"`
	FullName      *string   `json:"full_name"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// userDTOFromDomain конвертирует доменный объект User в DTO для HTTP-ответа.
func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:            user.ID,
		Version:       user.Version,
		Email:         user.Email,
		EmailVerified: user.EmailVerified,
		FullName:      user.FullName,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}

// usersDTOFromDomains конвертирует список доменных объектов в список DTO.
func usersDTOFromDomains(users []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(users))

	for i, user := range users {
		usersDTO[i] = userDTOFromDomain(user)
	}

	return usersDTO
}
