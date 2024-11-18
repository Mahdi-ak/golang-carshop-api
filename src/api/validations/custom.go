package validations

import (
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

type ValidationErrors struct {
	Property string `json:"property"`
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Massage  string `json:"massage"`
}

func GetValidationErrors(err error) *[]ValidationErrors {
	var validationErrors []ValidationErrors
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, err := range err.(validator.ValidationErrors) {
			var el ValidationErrors
			el.Property = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			el.Massage = err.Error()
			validationErrors = append(validationErrors, el)
		}
		return &validationErrors
	}
	return nil

}
