package helper

import "github.com/Mahdi-ak/golang-carshop-api/src/api/validations"

type BaseHttpResponse struct {
	Result           any                             `json:"result"`
	Success          bool                            `json:"success"`
	ResultCode       ResultCode                      `json:"resultCode"`
	ValidationErrors *[]validations.ValidationErrors `json:"validationErrors"`
	Error            any                             `json:"errors"`
}

func GenerateBaseResponse(result any, success bool, resultCode ResultCode) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Success:    success,
		ResultCode: resultCode,
	}
}

func GenerateBaseResponseWithError(result any, success bool, resultCode ResultCode, err error) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Success:    success,
		ResultCode: resultCode,
		Error:      err.Error(),
	}

}

func GenerateBaseResponseWithAnyError(result any, success bool, resultCode ResultCode, err any) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Success:    success,
		ResultCode: resultCode,
		Error:      err,
	}
}

func GenerateBaseResponseWithValidationError(result any, success bool, resultCode ResultCode, err error) *BaseHttpResponse {
	return &BaseHttpResponse{Result: result,
		Success:          success,
		ResultCode:       resultCode,
		ValidationErrors: validations.GetValidationErrors(err),
	}
}
