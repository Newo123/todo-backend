package domain

import "errors"

var (
	// ErrNotFound — запрашиваемая сущность не найдена (HTTP 404).
	ErrNotFound = errors.New("not found")

	// ErrInvalidArgument — переданы некорректные данные (HTTP 400).
	ErrInvalidArgument = errors.New("invalid argument")

	// ErrConflict — конфликт при обновлении, обычно из-за конкурентного
	// изменения той же записи (HTTP 409).
	ErrConflict = errors.New("conflict")

	// ErrUnauthorized - не авторизован (HTTP 401)
	ErrUnauthorized = errors.New("unauthorized")

	// ErrForbidden - доступ закрыт (HTTP 403)
	ErrForbidden = errors.New("forbidden")
)
