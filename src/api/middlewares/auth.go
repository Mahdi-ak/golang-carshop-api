package middlewares

import (
	"net/http"
	"strings"

	"github.com/Mahdi-ak/golang-carshop-api/src/api/helper"
	"github.com/Mahdi-ak/golang-carshop-api/src/config"
	"github.com/Mahdi-ak/golang-carshop-api/src/constants"
	service_errors "github.com/Mahdi-ak/golang-carshop-api/src/pkg/sevice_errors"
	"github.com/Mahdi-ak/golang-carshop-api/src/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthenticationMiddleware(cfg *config.Config) gin.HandlerFunc {
	var tokenService = services.NewTokenService(cfg)

	return func(c *gin.Context) {
		var err error
		claimMap := map[string]interface{}{}
		auth := c.GetHeader(constants.AuthorizationHeaderKey)
		token := strings.Split(auth, " ")
		if auth == "" {
			err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenRequired}
		} else {
			claimMap, err = tokenService.GetClaims(token[1])
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenExpired}
				default:
					err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenInvalid}
				}
			}
		}
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.GenerateBaseResponseWithError(nil, false, -2, err))
			return
		}
		c.Set(constants.UserIdKey, claimMap[constants.UserIdKey])
		c.Set(constants.UsernameKey, claimMap[constants.UsernameKey])
		c.Set(constants.FirstNameKey, claimMap[constants.FirstNameKey])
		c.Set(constants.LastNameKey, claimMap[constants.LastNameKey])
		c.Set(constants.EmailKey, claimMap[constants.EmailKey])
		c.Set(constants.MobileNumberKey, claimMap[constants.MobileNumberKey])
		c.Set(constants.RolesKey, claimMap[constants.RolesKey])
		c.Set(constants.ExpireTimeKey, claimMap[constants.ExpireTimeKey])

		c.Next()
	}
}

func AuthorizationMiddleware(validRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetStringSlice(constants.RolesKey)) == 0 {
			c.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, -3))
			return
		}
		rolesVal := c.Keys[constants.RolesKey]
		if rolesVal == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, -3))
			return
		}
		roles := rolesVal.([]interface{})
		val := map[string]int{}
		for _, item := range roles {
			val[item.(string)] = 0
		}

		for _, item := range validRoles {
			if _, ok := val[item]; ok {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, -3))
	}
}
