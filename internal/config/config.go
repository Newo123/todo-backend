package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config — Базовая конфиграция всего приложения
// Все поля читаются из переменных окружения с префиксом BASE_.
type Config struct {
	// TimeZone задаёт часовой пояс для time.Local во всём приложении.
	// Все time.Now() будут возвращать время в этом часовом поясе.
	TimeZone *time.Location `envconfig:"TIME_ZONE" default:"UTC"`
}

// NewConfig читает конфигурацию из переменных окружения.
// Переменная BASE_TIME_ZONE принимает IANA-идентификатор часового пояса:
//   - UTC
//   - Europe/Berlin
//   - America/New_York
//   - Europe/Moscow
func NewConfig() (Config, error) {
	var config Config
	if err := envconfig.Process("BASE", &config); err != nil {
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}

	return config, nil
}

// NewConfigMust — «Must»-вариант конструктора: паникует при ошибке.
// Используется при старте приложения, когда работа без конфигурации невозможна.
func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get Base config: %w", err)
		panic(err)
	}

	return config
}
