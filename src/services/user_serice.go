package services

import (
	"github.com/Mahdi-ak/golang-carshop-api/src/api/dto"
	"github.com/Mahdi-ak/golang-carshop-api/src/common"
	"github.com/Mahdi-ak/golang-carshop-api/src/config"
	"github.com/Mahdi-ak/golang-carshop-api/src/constants"
	"github.com/Mahdi-ak/golang-carshop-api/src/data/db"
	"github.com/Mahdi-ak/golang-carshop-api/src/data/models"
	"github.com/Mahdi-ak/golang-carshop-api/src/pkg/logging"
	service_errors "github.com/Mahdi-ak/golang-carshop-api/src/pkg/sevice_errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	logger       logging.Logger
	cfg          *config.Config
	otpService   *OtpService
	database     *gorm.DB
	tokenService *TokenService
}

func NewUserService(cfg *config.Config) *UserService {
	database := db.GetDb()
	logger := logging.NewLogger(cfg)
	return &UserService{
		logger:       logger,
		cfg:          cfg,
		otpService:   NewOtpService(cfg),
		database:     database,
		tokenService: NewTokenService(cfg),
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

func (s *UserService) existsByEmail(email string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists).
		Error; err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (s *UserService) existsByUsername(username string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("username = ?", username).
		Find(&exists).
		Error; err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (s *UserService) existsByMobileNumber(mobileNumber string) (bool, error) {
	var exists bool
	if err := s.database.Model(&models.User{}).
		Select("count(*) > 0").
		Where("mobile_number = ?", mobileNumber).
		Find(&exists).
		Error; err != nil {
		s.logger.Error(logging.Postgres, logging.Select, err.Error(), nil)
		return false, err
	}
	return exists, nil
}

func (s *UserService) getDefaultRole() (roleId int, err error) {

	if err = s.database.Model(&models.Role{}).
		Select("id").
		Where("name = ?", constants.DefaultRoleName).
		First(&roleId).Error; err != nil {
		return 0, err
	}
	return roleId, nil
}

// Register by username
func (s *UserService) RegisterByUsername(req *dto.RegisterUserByUsernameRequest) error {
	u := models.User{
		Username:  req.Username,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}
	exists, err := s.existsByEmail(req.Email)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.EmailExists}
	}
	exists, err = s.existsByUsername(req.Username)
	if err != nil {
		return err
	}
	if exists {
		return &service_errors.ServiceError{EndUserMessage: service_errors.UsernameExists}
	}

	bytePassword := []byte(req.Password)
	hashPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
		return err
	}
	u.Password = string(hashPassword)
	roleId, err := s.getDefaultRole()
	if err != nil {
		s.logger.Error(logging.Postgres, logging.DefaultRoleNotFound, err.Error(), nil)
		return err
	}

	tx := s.database.Begin()
	err = tx.Create(&u).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err
	}

	err = tx.Create(&models.UserRole{
		RoleId: roleId,
		UserId: u.Id,
	}).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
		return err
	}
	tx.Commit()
	return nil
}

// Register/login by mobile number
func (s *UserService) RegisterLoginByMobileNumber(req *dto.RegisterLoginByMobileRequest) (*dto.TokenDetail, error) {
	err := s.otpService.ValidateOtp(req.MobileNumber, req.Otp)
	if err != nil {
		return nil, err
	}

	exists, err := s.existsByMobileNumber(req.MobileNumber)
	if err != nil {
		return nil, err
	}

	u := models.User{MobileNumber: req.MobileNumber, Username: req.MobileNumber}

	if exists {
		var user models.User
		err = s.database.
			Model(&models.User{}).
			Where("username = ?", u.Username).
			Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
				return tx.Preload("Role")
			}).
			Find(&user).Error
		if err != nil {
			return nil, err
		}
		tdto := tokenDto{UserId: user.Id, FirstName: user.FirstName,
			LastName: user.LastName, Email: user.Email,
			mobileNumber: user.MobileNumber}

		if len(*user.UserRoles) > 0 {
			for _, ur := range *user.UserRoles {
				tdto.Roles = append(tdto.Roles, ur.Role.Name)
			}
		}

		token, err := s.tokenService.GenerateToken(&tdto)
		if err != nil {
			return nil, err
		}
		return token, nil

	} else {

		bytePassword := []byte(common.GeneratePassword())
		hashPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
		if err != nil {
			s.logger.Error(logging.General, logging.HashPassword, err.Error(), nil)
			return nil, err
		}
		u.Password = string(hashPassword)
		roleId, err := s.getDefaultRole()
		if err != nil {
			s.logger.Error(logging.Postgres, logging.DefaultRoleNotFound, err.Error(), nil)
			return nil, err
		}

		tx := s.database.Begin()
		err = tx.Create(&u).Error
		if err != nil {
			tx.Rollback()
			s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
			return nil, err
		}

		err = tx.Create(&models.UserRole{
			RoleId: roleId,
			UserId: u.Id,
		}).Error
		if err != nil {
			tx.Rollback()
			s.logger.Error(logging.Postgres, logging.Rollback, err.Error(), nil)
			return nil, err
		}
		tx.Commit()

		var user models.User
		err = s.database.
			Model(&models.User{}).
			Where("username = ?", u.Username).
			Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
				return tx.Preload("Role")
			}).
			Find(&user).Error
		if err != nil {
			return nil, err
		}
		tdto := tokenDto{UserId: user.Id, FirstName: user.FirstName,
			LastName: user.LastName, Email: user.Email,
			mobileNumber: user.MobileNumber}

		if len(*user.UserRoles) > 0 {
			for _, ur := range *user.UserRoles {
				tdto.Roles = append(tdto.Roles, ur.Role.Name)
			}
		}

		token, err := s.tokenService.GenerateToken(&tdto)
		if err != nil {
			return nil, err
		}
		return token, nil

	}

}

// login by username
func (s *UserService) LoginByUsername(req *dto.LoginByUsernameRequest) (*dto.TokenDetail, error) {
	var user models.User
	err := s.database.
		Model(&models.User{}).
		Where("username = ?", req.Username).
		Preload("UserRoles", func(tx *gorm.DB) *gorm.DB {
			return tx.Preload("Role")
		}).
		Find(&user).Error
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	tdto := tokenDto{UserId: user.Id, FirstName: user.FirstName,
		LastName: user.LastName, Email: user.Email,
		mobileNumber: user.MobileNumber}

	if len(*user.UserRoles) > 0 {
		for _, ur := range *user.UserRoles {
			tdto.Roles = append(tdto.Roles, ur.Role.Name)
		}
	}

	token, err := s.tokenService.GenerateToken(&tdto)
	if err != nil {
		return nil, err
	}
	return token, nil
}
