package errorWrapper

const (
	InternalError         = "internal_error"
	JsonParsingError      = "json_parsing_error"
	OrderCreateValidation = "order_create_validation"
	OrderCreateError      = "order_create_error"
)

type Error struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

func NewError(t string, message string) *Error {
	return &Error{
		Type:    t,
		Message: message,
	}
}
