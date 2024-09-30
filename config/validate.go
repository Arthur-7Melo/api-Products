package config

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/pt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ptbr_translations "github.com/go-playground/validator/v10/translations/pt_BR"
)

var (
	Validate = validator.New()
	transl ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		pt := pt.New()
		unt := ut.New(pt, pt)
		transl, _ = unt.GetTranslator("pt")
		ptbr_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateProductError(
	validation_err error,
) *ProductError{
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationErr validator.ValidationErrors

	if errors.As(validation_err, &jsonErr){
		return NewBadRequestError("Campo Inválido")
	} else if errors.As(validation_err, &jsonValidationErr) {
		errorCauses := []Causes{}
			for _, e := range jsonValidationErr {
				cause := Causes{
					Field: e.Field(),
					Message: e.Translate(transl),
				}
				errorCauses = append(errorCauses, cause)						
			}
		return NewBadRequestValidationError("Alguns campos são inválidos!", errorCauses)
	} 
	return NewInternalServerError("Erro ao tentar converter campos")
}