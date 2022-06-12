package database

import (
	"dousheng/config"
	"github.com/go-redis/redis"
	"log"
)

var redisDB *redis.Client

func InitRedisClient(cfg *config.Config) error {
	redisCfg := cfg.Redis
	redisDB := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Address,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
		PoolSize: redisCfg.PoolSize,
	})
	res, err := redisDB.Ping().Result()
	if err != nil {
		return err
	}
	log.Printf("[my-log] Ping Reids Client : " + res)
	return nil
}
