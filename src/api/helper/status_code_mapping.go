package helper

import (
	"net/http"

	service_errors "github.com/Mahdi-ak/golang-carshop-api/src/pkg/sevice_errors"
)

var StatusCodeMapping = map[string]int{

	// otp
	service_errors.OtpExists:   409,
	service_errors.OtpUsed:     409,
	service_errors.OtpNotValid: 400,
}

func TranslateErrorToStatusCode(err error) int {

	value, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}
	return value

}
