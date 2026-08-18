package response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/Newo123/todo-backend/internal/domain"
	"github.com/Newo123/todo-backend/internal/infrastructure/logger"
)

// HTTPResponseHandler инкапсулирует логику записи HTTP-ответов.
// Хранит логгер и ResponseWriter, чтобы обработчикам не нужно было
// передавать их каждый раз явно.
type Response struct {
	log logger.Logger
	rw  http.ResponseWriter
}

func NewResponse(log logger.Logger, rw http.ResponseWriter) *Response {
	return &Response{
		log: log,
		rw:  rw,
	}
}

// JSON сериализует responseBody в JSON и записывает в ответ с указанным статус-кодом.
// Content-Type автоматически определяется json.NewEncoder.
func (h *Response) JSON(statusCode int, body any) {
	h.rw.WriteHeader(statusCode)

	if err := json.NewEncoder(h.rw).Encode(body); err != nil {
		h.log.Error("write HTTP response", logger.Error(err))
	}
}

// OK обертка над методом JSON со статусом 200
func (h *Response) OK(body any) {
	h.JSON(http.StatusOK, body)
}

// NoContentResponse отправляет HTTP 204 No Content — используется при успешном DELETE.
func (h *Response) NoContent() {
	h.rw.WriteHeader(http.StatusNoContent)
}

// ErrorResponse транслирует core ошибку в HTTP-статус через errors.Is().
//
// Маппинг:
//   - ErrInvalidArgument → 400
//   - ErrNotFound        → 404
//   - ErrConflict        → 409
//   - остальное          → 500
//
// Каждый тип ошибки логируется на соответствующем уровне (Warn/Debug/Error).
func (h *Response) Error(err error, msg string) {
	var (
		statusCode int
		logFunc    func(string, ...logger.Field)
	)

	switch {
	case errors.Is(err, domain.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = h.log.Warn
	case errors.Is(err, domain.ErrNotFound):
		statusCode = http.StatusNotFound
		logFunc = h.log.Debug
	case errors.Is(err, domain.ErrConflict):
		statusCode = http.StatusConflict
		logFunc = h.log.Warn
	case errors.Is(err, domain.ErrUnauthorized):
		statusCode = http.StatusUnauthorized
		logFunc = h.log.Warn
	case errors.Is(err, domain.ErrForbidden):
		statusCode = http.StatusForbidden
		logFunc = h.log.Warn
	default:
		statusCode = http.StatusInternalServerError
		logFunc = h.log.Error
	}

	logFunc(msg, logger.Error(err))

	h.errorResponse(statusCode, err, msg)
}

// PanicResponse формирует HTTP 500 при перехвате паники.
// Вызывается из middleware Recovery — см. internal/transport/http/middleware/recovery.go.
func (h *Response) Panic(p any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panic: %v", p)

	h.log.Error(msg, logger.Error(err))

	h.errorResponse(statusCode, err, msg)
}

// errorResponse — внутренний метод: собирает ErrorResponse и вызывает JSON
func (h *Response) errorResponse(statusCode int, err error, msg string) {
	response := ErrorResponse{
		Error:   err.Error(),
		Message: msg,
	}

	h.JSON(statusCode, response)
}
