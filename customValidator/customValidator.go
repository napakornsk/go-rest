package customevalidator

import (
	"fmt"
	"unicode"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	Validator *validator.Validate
}

func InitValidator(v *validator.Validate) *Validator {
	return &Validator{
		Validator: v,
	}
}

func (Validator) Password(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}

func (Validator) GetErrorMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters long", fe.Field(), fe.Param())
	case "alphanum":
		return fmt.Sprintf("%s must only contain alphanumeric characters", fe.Field())
	case "password":
		return fmt.Sprintf("%s must only contain atleast one of special symbol, numeric, uppercase and lowercase", fe.Field())
	default:
		return fmt.Sprintf("%s is invalid", fe.Field())
	}
}
