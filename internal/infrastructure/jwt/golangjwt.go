package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type golangJWT struct {
	secret          []byte
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
	issuer          string
}

type claims struct {
	UserID    uuid.UUID `json:"sub"`
	TokenType TokenType `json:"type"`

	jwt.RegisteredClaims
}

func NewGolangJWT(config Config) JWTManager {
	return &golangJWT{
		secret:          []byte(config.Secret),
		accessTokenTTL:  config.AccessTokenTTL,
		refreshTokenTTL: config.RefreshTokenTTL,
		issuer:          config.Issuer,
	}
}

func (j *golangJWT) Generate(userID uuid.UUID, tokenType TokenType) (string, error) {
	if userID == uuid.Nil {
		return "", ErrInvalidUserID
	}

	var expiration time.Duration

	switch tokenType {
	case TokenTypeAccess:
		expiration = j.accessTokenTTL

	case TokenTypeRefresh:
		expiration = j.refreshTokenTTL

	default:
		return "", ErrInvalidTokenType
	}

	now := time.Now()

	c := claims{
		UserID:    userID,
		TokenType: tokenType,

		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   userID.String(),
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(expiration)),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		c,
	)

	signedToken, err := token.SignedString(j.secret)
	if err != nil {
		return "", fmt.Errorf("%w: %w", ErrTokenGeneration, err)
	}

	return signedToken, nil
}

func (j *golangJWT) Validate(tokenString string, tokenType TokenType) (*Claims, error) {
	if tokenString == "" {
		return nil, ErrInvalidToken
	}

	if tokenType != TokenTypeAccess &&
		tokenType != TokenTypeRefresh {
		return nil, ErrInvalidTokenType
	}

	var c claims

	token, err := jwt.ParseWithClaims(
		tokenString,
		&c,
		func(token *jwt.Token) (any, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, ErrInvalidSigningMethod
			}

			return j.secret, nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrInvalidToken, err)
	}

	if !token.Valid {
		return nil, ErrInvalidToken
	}

	if c.TokenType != tokenType {
		return nil, ErrInvalidToken
	}

	userID, err := uuid.Parse(c.Subject)
	if err != nil {
		return nil, fmt.Errorf(
			"%w: invalid user id: %w",
			ErrInvalidToken,
			err,
		)
	}

	return &Claims{
		UserID:    userID,
		TokenType: c.TokenType,
	}, nil
}
