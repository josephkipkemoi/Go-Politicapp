package utils

import (
	"log"

	"github.com/go-playground/validator/v10"
)

// ValidateRequest returns found errors stored in a slice and true if errors are found
// empty slice and false otherwise
func ValidateRequest(err error) ([]string, bool) {
	errs := []string{}

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return errs, false
		}
		for _, e := range err.(validator.ValidationErrors) {
			log.Println(e)
			errs = append(errs, e.Field()+"field is required")
		}
		return errs, false
	}
	return errs, true
}
