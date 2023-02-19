package output

type ErrorMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Errors []ErrorMessage `json:"errors"`
	Code   int            `json:"code"`
}

func (e *ErrorResponse) AddError(field string, message string) {
	e.Errors = append(e.Errors, ErrorMessage{Message: message, Field: field})
}

func NewErrorResponse(code int) *ErrorResponse {
	return &ErrorResponse{Errors: []ErrorMessage{}, Code: code}
}
