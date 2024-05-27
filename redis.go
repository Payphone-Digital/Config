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

// func SetRedisClientCluster() *redis.Client {
// redis.NewClusterClient(&redis.ClusterOptions{
// 	Addrs: []string{"localhost:6379", "localhost:6380", "localhost:6381"},
// 	Password: config.Addr, // Ganti dengan alamat node Redis Cluster
// })
// }
