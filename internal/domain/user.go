package domain

import (
	"fmt"
	"net/mail"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID
	Version       int64
	Email         string
	EmailVerified bool
	PasswordHash  string
	FullName      *string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewUser(
	id uuid.UUID,
	version int64,
	email string,
	emailVerified bool,
	passwordHash string,
	fullName *string,
	createdAt time.Time,
	updatedAt time.Time,
) User {
	return User{
		ID:            id,
		Version:       version,
		Email:         email,
		EmailVerified: emailVerified,
		PasswordHash:  passwordHash,
		FullName:      fullName,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}

func CreateUser(
	email string,
	passwordHash string,
	fullName *string,
	emailVerified bool,
) (User, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return User{}, fmt.Errorf("create uuid: %w", err)
	}

	now := time.Now()

	user := NewUser(
		id,
		int64(1),
		email,
		emailVerified,
		passwordHash,
		fullName,
		now,
		now,
	)

	if err := user.validate(); err != nil {
		return User{}, fmt.Errorf("validate user domain: %w", err)
	}

	return user, nil
}

func (u *User) validate() error {
	email := strings.TrimSpace(u.Email)
	if email == "" {
		return fmt.Errorf("email is required: %w", ErrInvalidArgument)
	}

	if len([]rune(email)) > 254 {
		return fmt.Errorf("email is too long: %w", ErrInvalidArgument)
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return fmt.Errorf("invalid email: %w: %v", ErrInvalidArgument, err)
	}

	if u.PasswordHash == "" {
		return fmt.Errorf("password hash is required: %w", ErrInvalidArgument)
	}

	if u.FullName != nil {
		fullName := strings.TrimSpace(*u.FullName)
		fullNameLen := utf8.RuneCountInString(fullName)

		if fullNameLen < 3 || fullNameLen > 100 {
			return fmt.Errorf(
				"invalid `FullName` len: %d: %w",
				fullNameLen,
				ErrInvalidArgument,
			)
		}
	}

	return nil
}
