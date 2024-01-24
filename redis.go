package config

import "github.com/go-redis/redis/v8"

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

// SetRedisClient allows dynamic configuration of the Redis client
func SetRedisClient(config RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
}
