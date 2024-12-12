package services

import (
	"fmt"
	"time"

	"github.com/Mahdi-ak/golang-carshop-api/src/config"
	"github.com/Mahdi-ak/golang-carshop-api/src/data/cache"
	"github.com/Mahdi-ak/golang-carshop-api/src/pkg/logging"
	service_errors "github.com/Mahdi-ak/golang-carshop-api/src/pkg/sevice_errors"
	"github.com/redis/go-redis/v9"
)

type OtpService struct {
	logger logging.Logger
	cfg    *config.Config
	redis  *redis.Client
}

// OtpData is a struct for otp data
type OtpDto struct {
	Value string
	Used  bool
}

func NewOtpService(cfg *config.Config) *OtpService {
	logger := logging.NewLogger(cfg)
	redisClient := cache.GetRedis()
	return &OtpService{logger: logger, cfg: cfg, redis: redisClient}
}

func (s *OtpService) SetOtp(mobileNumber string, otp string) error {
	key := fmt.Sprintf("otp:%s", mobileNumber)
	val := &OtpDto{
		Value: otp,
		Used:  false,
	}

	res, err := cache.Get[OtpDto](s.redis, key)
	if err == nil && !res.Used {
		return service_errors.ServiceError{EndUserMessage: "Otp exists"}
	} else if err == nil && res.Used {
		return service_errors.ServiceError{EndUserMessage: "Otp Used"}
	}
	err = cache.Set(s.redis, key, val, s.cfg.Otp.ExpireTime*time.Second)
	if err != nil {
		return err
	}
	return nil
}

func (s *OtpService) ValidateOtp(mobileNumber string, otp string) error {

	key := fmt.Sprintf("otp:%s", mobileNumber)
	res, err := cache.Get[OtpDto](s.redis, key)

	if err != nil {
		return err
	} else if res.Used {
		return service_errors.ServiceError{EndUserMessage: "otp used"}
	} else if !res.Used && res.Value != otp {
		return service_errors.ServiceError{EndUserMessage: "otp invalid"}
	} else if !res.Used && res.Value == otp {
		res.Used = true
		err = cache.Set(s.redis, key, res, s.cfg.Otp.ExpireTime*time.Second)
		if err != nil {
			return err
		}
	}
	return nil
}
