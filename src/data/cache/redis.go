package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Mahdi-ak/golang-carshop-api/src/config"
	"github.com/Mahdi-ak/golang-carshop-api/src/pkg/logging"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client
var ctx = context.Background()
var log = logging.NewLogger(config.GetConfig())

// InitRedis initializes the redis client
func InitRedis(cfg *config.Config) error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:         cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.Db,
		DialTimeout:  cfg.Redis.DialTimeout * time.Second,
		ReadTimeout:  cfg.Redis.ReadTimeout * time.Second,
		WriteTimeout: cfg.Redis.WriteTimeout * time.Second,
		PoolSize:     cfg.Redis.PoolSize,
		PoolTimeout:  cfg.Redis.PoolTimeout * time.Second,
		// MinIdleConns: 10,
	})

	ctx := context.Background()
	if res := redisClient.Ping(ctx).String(); res != "ping: PONG" {
		log.Error(logging.Redis, logging.Startup, "Redis connection failed", nil)
	}

	return nil
}

func Set[T any](c *redis.Client, key string, value T, expiration time.Duration) error {
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}

	err = c.Set(ctx, key, v, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func Get[T any](c *redis.Client, key string) (T, error) {
	var res = *new(T)
	v, err := c.Get(ctx, key).Result()
	if err != nil {
		return res, err
	}

	err = json.Unmarshal([]byte(v), &res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func GetRedis() *redis.Client {
	return redisClient

}
func CloseRedis() {
	redisClient.Close()
}
