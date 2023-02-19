package translation

import (
	"errors"
	"github.com/gin-gonic/gin/binding"
	english_locales "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	english_translations "github.com/go-playground/validator/v10/translations/en"
)

var Translator ut.Translator

func RegisterValidationTranslations() error {
	eng := english_locales.New()
	uni := ut.New(eng, eng)
	Translator, _ = uni.GetTranslator("en")

	validatorValidate := binding.Validator.Engine().(*validator.Validate)

	if err := english_translations.RegisterDefaultTranslations(validatorValidate, Translator); err != nil {
		return errors.New("could not register translations for `en`")
	}

	return nil
}
