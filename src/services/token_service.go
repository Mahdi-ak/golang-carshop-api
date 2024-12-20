package services

import (
	"time"

	"github.com/Mahdi-ak/golang-carshop-api/src/api/dto"
	"github.com/Mahdi-ak/golang-carshop-api/src/config"
	"github.com/Mahdi-ak/golang-carshop-api/src/pkg/logging"
	service_errors "github.com/Mahdi-ak/golang-carshop-api/src/pkg/sevice_errors"
	"github.com/golang-jwt/jwt"
)

type TokenService struct {
	logger logging.Logger
	cfg    *config.Config
}

type tokenDto struct {
	UserId       int
	FirstName    string
	LastName     string
	Email        string
	Roles        []string
	mobileNumber string
	UserName     string
}

func NewTokenService(cfg *config.Config) *TokenService {
	logger := logging.NewLogger(cfg)
	return &TokenService{logger: logger, cfg: cfg}
}

func (s *TokenService) GenerateToken(token *tokenDto) (*dto.TokenDetail, error) {
	tokenDetail := &dto.TokenDetail{
		AccessTokenExpireTime:  time.Now().Add(s.cfg.JWT.AccessTokenExpireDuration * time.Minute).Unix(),
		RefreshTokenExpireTime: time.Now().Add(s.cfg.JWT.RefreshTokenExpireDuration).Unix(),
	}

	AccessTokenClaims := jwt.MapClaims{
		"user_id":    token.UserId,
		"first_name": token.FirstName,
		"last_name":  token.LastName,
		"email":      token.Email,
		"role":       token.Roles,
		"user_name":  token.UserName,
		"mobile":     token.mobileNumber,
		"ex":         tokenDetail.AccessTokenExpireTime,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodES256, AccessTokenClaims)

	var err error

	tokenDetail.AccessToken, err = at.SignedString([]byte(s.cfg.JWT.Secret))
	if err != nil {
		return nil, err
	}

	RefreshTokenClaims := jwt.MapClaims{
		"user_id": token.UserId,
		"ex":      tokenDetail.RefreshTokenExpireTime,
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodES256, RefreshTokenClaims)
	tokenDetail.RefreshToken, err = rt.SignedString([]byte(s.cfg.JWT.Secret))
	if err != nil {
		return nil, err
	}

	return tokenDetail, nil

}

func (s *TokenService) VerifyToken(token string) (*jwt.Token, error) {

	tk, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, &service_errors.ServiceError{EndUserMessage: service_errors.UnExpectedError}
		}
		return []byte(s.cfg.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return tk, nil
}

func (s *TokenService) GetClaims(token string) (claimMap map[string]interface{}, err error) {
	claimMap = make(map[string]interface{})

	verifyToken, err := s.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	claims, ok := verifyToken.Claims.(jwt.MapClaims)
	if ok && verifyToken.Valid {
		for key, value := range claims {
			claimMap[key] = value
		}
		return claimMap, nil
	}
	return nil, &service_errors.ServiceError{EndUserMessage: service_errors.ClaimsNotFound}
}
