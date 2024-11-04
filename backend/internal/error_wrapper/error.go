package errorWrapper

import "fmt"

const (
	InternalError         = "internal_error"
	JsonParsingError      = "json_parsing_error"
	OrderCreateValidation = "order_create_validation"
	OrderCreateError      = "order_create_error"
	ProductCountError     = "product_count_error"
	JwtAuthMiddleware     = "jwt_auth_middleware"
)

type Error struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func (err *Error) Error() string {
	return fmt.Sprintf("%s: %s", err.Type, err.Message)
}

func NewError(t string, message string) *Error {
	return &Error{
		Type:    t,
		Message: message,
	}
}
