package validators

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidateUsreFisrtName(field validator.FieldLevel) bool {
	firstName := field.Field().String()
	re := regexp.MustCompile(`^[A-Za-zА-Яа-яЁё\s]*$`)
	return re.MatchString(firstName)
}
