package validation

import (
	"github.com/go-playground/validator/v10"
	"time"
)

// registerValidators sets up custom validators on application start
func RegisterValidators(validate *validator.Validate) {

	err := validate.RegisterValidation("rfc3339", func(fl validator.FieldLevel) bool {
		if fl.Field().Type().Name() == "string" {
			_, err := time.Parse(time.RFC3339, fl.Field().String())
			return err == nil
		}
		return false
	})
	if err != nil {
		return
	}
}