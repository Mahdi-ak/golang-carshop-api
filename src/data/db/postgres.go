package db

import (
	"fmt"
	"time"

	"github.com/Mahdi-ak/golang-carshop-api/src/config"
	"github.com/Mahdi-ak/golang-carshop-api/src/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB
var logger = logging.NewLogger(config.GetConfig())

func InitDb(cfg *config.Config) error {
	var err error

	cnn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DbName, cfg.Postgres.SSLMode)

	dbClient, err = gorm.Open(postgres.Open(cnn), &gorm.Config{})
	if err != err {
		return err
	}

	sqlDb, _ := dbClient.DB()
	err = sqlDb.Ping()
	if err != err {
		return err
	}

	sqlDb.SetConnMaxIdleTime(time.Duration(cfg.Postgres.MaxIdleConns))
	sqlDb.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)

	logger.Info(logging.Postgres, logging.Startup, "Database connection established", nil)
	return nil
}
func GetDb() *gorm.DB {
	return dbClient
}

func CloseDb() {
	con, _ := dbClient.DB()
	con.Close()
}
