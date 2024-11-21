package config

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
	Cors     CorsConfig
	Password PasswordConfig
	Otp      OtpConfig

	// Logger LogConfig
}

type ServerConfig struct {
	Port    int
	RunMode string
}

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
	SSLMode  string
}
type RedisConfig struct {
	Host               string
	Port               int
	Password           string
	Db                 int
	MinIdleConnections int
	PoolSize           int
	PoolTimeout        int
}

type PasswordConfig struct {
	IncludeChars     bool
	IncludeDigits    bool
	MinLength        int
	MaxLength        int
	IncludeUppercase bool
	IncludeLowercase bool
}
type CorsConfig struct {
	AllowOrigins string
}

type OtpConfig struct {
	ExpireTime time.Duration
	Digits     int
	Limiter    time.Duration
}

// GetConfig retrieves the application configuration settings.
// It first checks the environment variable APP_ENV to determine the config file to use.
// If the file is found, it is read into a viper.Viper object.
// The object is then unmarshaled into a Config object, which is returned.
// If the config file is not found, an error is logged and returned.

func GetConfig() *Config {
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))

	v, err := LogConfig(cfgPath, "yml")
	if err != nil {
		log.Fatalf("Error in load config %v", err)
	}

	cfg, err := ParseConfig(v)
	if err != nil {
		log.Fatalf("Error in pars config %v", err)
	}
	return cfg

}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}
	return &cfg, nil

}
func LogConfig(filename string, fileType string) (*viper.Viper, error) {

	v := viper.New()
	v.SetConfigType(fileType)
	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	err := v.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("unable to read config, error: %v", err)
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil
}
func getConfigPath(env string) string {
	if env == "docker" {
		return "src/config/config-docker"
	} else if env == "production" {
		return "config/config-production"
	} else {
		return "src/config/config-development"
	}
}
