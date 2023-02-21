package translation

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	english_locales "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	english_translations "github.com/go-playground/validator/v10/translations/en"
)

var Translator ut.Translator

const en_locale = "en"

func RegisterValidationTranslations() error {
	eng := english_locales.New()
	uni := ut.New(eng, eng)
	Translator, _ = uni.GetTranslator(en_locale)

	validatorValidate := binding.Validator.Engine().(*validator.Validate)

	if err := english_translations.RegisterDefaultTranslations(validatorValidate, Translator); err != nil {
		return fmt.Errorf("could not register translations for `%s`", en_locale)
	}

	return nil
}
