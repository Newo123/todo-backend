package jwt

import "github.com/google/uuid"

type TokenType string

const (
	TokenTypeAccess  TokenType = "access"
	TokenTypeRefresh TokenType = "refresh"
)

type JWTManager interface {
	Generate(userID uuid.UUID, tokenType TokenType) (string, error)
	Validate(token string, tokenType TokenType) (*Claims, error)
}

type Claims struct {
	UserID    uuid.UUID
	TokenType TokenType
}
