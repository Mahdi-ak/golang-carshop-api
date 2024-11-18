package validations

import (
	"github.com/Mahdi-ak/golang-carshop-api/src/common"
	"github.com/go-playground/validator/v10"
)

// mobile number validation
func IranianMobileNumberValidator(fld validator.FieldLevel) bool {

	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}

	return common.IranianMobileNumberValidate(value)
}
