package services

import (
	"github.com/Mahdi-ak/golang-carshop-api/src/api/dto"
	"github.com/Mahdi-ak/golang-carshop-api/src/common"
	"github.com/Mahdi-ak/golang-carshop-api/src/config"
	"github.com/Mahdi-ak/golang-carshop-api/src/data/db"
	"github.com/Mahdi-ak/golang-carshop-api/src/pkg/logging"
	"gorm.io/gorm"
)

type UserService struct {
	logger     logging.Logger
	cfg        *config.Config
	otpService *OtpService
	database   *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	database := db.GetDb()
	logger := logging.NewLogger(cfg)
	return &UserService{
		logger:     logger,
		cfg:        cfg,
		otpService: NewOtpService(cfg),
		database:   database,
	}

}

func (s *UserService) SendOtp(req *dto.GetOtpRequest) error {
	otp := common.GenerateOtp()
	err := s.otpService.SetOtp(req.MobileNumber, otp)
	if err != nil {
		return err
	}
	return nil
}
