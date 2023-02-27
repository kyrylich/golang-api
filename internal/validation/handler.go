package validation

import (
	"errors"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"golangpet/internal/dto/output"
	"net/http"
)

func CreateValidationResponse(translator ut.Translator, validationErr error) output.ErrorResponse {
	var ve validator.ValidationErrors

	if errors.As(validationErr, &ve) {
		out := make([]output.ErrorMessage, len(ve))
		for i, fe := range ve {
			out[i] = output.ErrorMessage{Field: fe.Field(), Message: fe.Translate(translator)}
		}
		return output.ErrorResponse{Errors: out, Code: http.StatusBadRequest}
	}

	return output.ErrorResponse{}
}
