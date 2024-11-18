package validations

import (
	"github.com/Mahdi-ak/golang-carshop-api/src/common"
	"github.com/go-playground/validator/v10"
)

// password validator
func PasswordValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		fld.Param()
		return false
	}

	return common.CheckPassword(value)
}
