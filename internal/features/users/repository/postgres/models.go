package postgres

import (
	"time"

	"github.com/Newo123/todo-backend/internal/domain"
	"github.com/Newo123/todo-backend/internal/infrastructure/postgres"
	"github.com/google/uuid"
)

// UserModel — структура для маппинга строки таблицы `todoapp.users` в Go-тип.
// Порядок полей совпадает с порядком столбцов в SELECT-запросах репозитория.
type UserModel struct {
	ID            uuid.UUID
	Version       int64
	Email         string
	EmailVerified bool
	PasswordHash  string
	FullName      *string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// Scan заполняет поля модели из результата запроса к БД.
func (m *UserModel) Scan(row postgres.Row) error {
	return row.Scan(
		&m.ID,
		&m.Version,
		&m.Email,
		&m.EmailVerified,
		&m.PasswordHash,
		&m.FullName,
		&m.CreatedAt,
		&m.UpdatedAt,
	)
}

// modelToDomain конвертирует модель БД в доменный объект.
func modelToDomain(model UserModel) domain.User {
	return domain.NewUser(
		model.ID,
		model.Version,
		model.Email,
		model.EmailVerified,
		model.PasswordHash,
		model.FullName,
		model.CreatedAt,
		model.UpdatedAt,
	)
}

// modelsToDomains конвертирует список моделей БД в список доменных объектов.
func modelsToDomains(models []UserModel) []domain.User {
	domains := make([]domain.User, len(models))

	for i, model := range models {
		domains[i] = modelToDomain(model)
	}

	return domains
}
