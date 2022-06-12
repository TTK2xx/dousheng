package database

import (
	"dousheng/config"
	"github.com/go-redis/redis"
	"log"
	"strconv"
)

var RedisDB *redis.Client

const SPLIT = ":"
const PREFIX_FOLLOW = "follow"
const PREFIX_FOLLOWER = "follower"

func InitRedisClient(cfg *config.Config) error {
	redisCfg := cfg.Redis
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Address,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
		PoolSize: redisCfg.PoolSize,
	})
	res, err := RedisDB.Ping().Result()
	if err != nil {
		return err
	}
	log.Printf("[my-log] Ping Reids Client : " + res)
	return nil
}

func GetFollowKey(userId int64) string {
	return strconv.FormatInt(userId, 10) + SPLIT + PREFIX_FOLLOW
}

func GetFollowerKey(userId int64) string {
	return strconv.FormatInt(userId, 10) + SPLIT + PREFIX_FOLLOWER
}
