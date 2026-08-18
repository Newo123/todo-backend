package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/Newo123/todo-backend/docs"
	"github.com/Newo123/todo-backend/internal/config"
	usersRepository "github.com/Newo123/todo-backend/internal/features/users/repository/postgres"
	usersService "github.com/Newo123/todo-backend/internal/features/users/service"
	usersHTTPHandler "github.com/Newo123/todo-backend/internal/features/users/transport/http"
	"github.com/Newo123/todo-backend/internal/infrastructure/hasher"
	"github.com/Newo123/todo-backend/internal/infrastructure/logger"
	"github.com/Newo123/todo-backend/internal/infrastructure/logger/zap"
	"github.com/Newo123/todo-backend/internal/infrastructure/postgres/pgx"
	"github.com/Newo123/todo-backend/internal/infrastructure/redis/goredis"
	"github.com/Newo123/todo-backend/internal/transport/http/middleware"
	"github.com/Newo123/todo-backend/internal/transport/http/server"
	"github.com/alexedwards/argon2id"
)

// Аннотации для автогенерации Swagger-документации (swaggo/swag).
// @title        Golang Todo API
// @version      1.0
// @description  Todo Application REST-API scheme
// @host         127.0.0.1:5050
// @BasePath     /api/v1
func main() {
	// Загружаем общую конфигурацию приложения
	// NewConfigMust — паттерн «Must»: паникует при ошибке, т.к. на старте
	// приложение не может продолжать работу с невалидной конфигурацией.
	config := config.NewConfigMust()
	time.Local = config.TimeZone

	// Создаём корневой контекст, который отменяется при получении SIGINT/SIGTERM
	// (Ctrl+C или команда `kill`). Это основа для graceful shutdown.
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	// Инициализируем логгер приложения
	// Пишет одновременно в stdout и в файл (см. internal/infrastructure/logger).
	log, err := zap.NewLogger(zap.NewConfigMust())
	if err != nil {
		fmt.Println("failed to init application logger:", err)
		os.Exit(1)
	}
	defer log.Close()

	log.Debug("application time zone", logger.Any("zone", time.Local))

	// Создаём пулл соединений с PostgreSQL через библиотеку pgx.
	// Пул переиспользует соединения, что гораздо эффективнее,
	// чем открывать новое соединение на каждый SQL запрос.
	postgresPool, err := pgx.NewPool(ctx, pgx.NewConfigMust())
	if err != nil {
		log.Fatal("failed to init postgres connection pool", logger.Error(err))
	}
	defer postgresPool.Close()
	log.Debug("initializing postgres connection pool")

	redisPool, err := goredis.NewPool(ctx, goredis.NewConfigMust())
	if err != nil {
		log.Fatal("failed to init redis connection pool", logger.Error(err))
	}
	defer redisPool.Close()
	log.Debug("initializing redis connection pool")

	// Ручное внедрение зависимостей (Dependency Injection):
	// Repository → Service → HTTP Handler.
	// Каждый слой знает только об интерфейсе нижележащего.
	// Это обеспечивает слабую связанность (loose coupling) и тестируемость.

	// Infra helpers

	hasher := hasher.NewArgon2IDHasher(argon2id.DefaultParams)

	// Repositories

	usersRepository := usersRepository.NewRepository(postgresPool)

	// Services

	usersService := usersService.NewService(usersRepository, hasher)

	// HTTP Handlers

	usersHTTPHandler := usersHTTPHandler.NewHTTPTransport(usersService)

	// Собираем HTTP-сервер с цепочкой middleware.
	// Middleware применяются ко всем маршрутам (Route) в порядке объявления:
	// CORS → RequestID → Logger → Trace → Panic recovery.
	log.Debug("initializing HTTP server")
	httpConfig := server.NewConfigMust()
	httpServer := server.NewHTTPServer(
		httpConfig,
		log,
		middleware.CORS(httpConfig.AllowedOrigins),
		middleware.RequestID(),
		middleware.Logger(log),
		middleware.Trace(),
		middleware.Recovery(),
	)

	// Регистрируем маршруты API v1.
	// APIVersionRouter автоматически добавляет префикс /api/v1 ко всем путям.
	routerV1 := server.NewRouter(server.ApiVersion1)
	routerV1.AddRoutes(usersHTTPHandler.Routes()...)

	httpServer.RegisterAPIRouters(routerV1)

	// Регистрируем Swagger UI по адресу /swagger/.
	httpServer.RegisterSwagger()

	// Запускаем сервер. Блокируется до получения сигнала завершения.
	// После сигнала выполняет graceful shutdown: ждёт завершения активных запросов.
	if err := httpServer.Run(ctx); err != nil {
		log.Error("HTTP server run error", logger.Error(err))
	}
}
