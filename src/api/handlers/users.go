package handlers

import (
	"net/http"

	"github.com/Mahdi-ak/golang-carshop-api/src/api/dto"
	"github.com/Mahdi-ak/golang-carshop-api/src/api/helper"
	"github.com/Mahdi-ak/golang-carshop-api/src/config"
	"github.com/Mahdi-ak/golang-carshop-api/src/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(cfg *config.Config) *UserHandler {
	service := services.NewUserService(cfg)
	return &UserHandler{service: service}
}

// SendOtp sends an OTP to the specified mobile number
//
// @Summary Send OTP
// @Description Sends an OTP to the specified mobile number
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.GetOtpRequest true "GetOtpRequest"
// @Success 201 {object} helper.BaseHttpResponse "OTP Sent"
// @Failure 400 {object} helper.BaseHttpResponse "Invalid Request"
// @Failure 500 {object} helper.BaseHttpResponse "Internal Server Error"
// @Router /v1/users/send-otp [post]
func (h *UserHandler) SendOtp(c *gin.Context) {
	req := new(dto.GetOtpRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		// bad request
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	// call the service
	err = h.service.SendOtp(req)
	if err != nil {
		// internal server error
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithAnyError(nil, false, -1, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("OTP Sent", true, 0))
}
