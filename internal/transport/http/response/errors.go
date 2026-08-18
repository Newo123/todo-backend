package response

// ErrorResponse — стандартная структура тела ответа при ошибке.
//   - Error   — полный текст ошибки (цепочка от обработчика до причины)
//   - Message — краткое человекочитаемое сообщение (что пытался сделать обработчик)
type ErrorResponse struct {
	Error   string  `json:"error"   example:"full error text"`
	Message string  `json:"message" example:"short human-readable message"`
	Code    ErrCode `json:"code" example:"error code NOT_FOUND"`
}

type ErrCode string

const (
	ErrCodeNotFound        ErrCode = "NOT_FOUND"
	ErrCodeAlreadyExists   ErrCode = "ALREADY_EXISTS"
	ErrCodeInvalidArgument ErrCode = "INVALID_ARGUMENT"
	ErrCodeConflict        ErrCode = "CONFLICT"
	ErrCodeUnauthorized    ErrCode = "UNAUTHORIZED"
	ErrCodeForbidden       ErrCode = "FORBIDDEN"
	ErrCodeInternal        ErrCode = "INTERNAL"
)
