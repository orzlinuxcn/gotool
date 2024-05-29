package redis

import (
	"github.com/go-redis/redis"
	"github.com/orzlinuxcn/gotool/consts"
)

var redisClient *redis.Client

func GetRedis() *redis.Client {
	if redisClient == nil {
		redisClient = InitClient()
	}
	return redisClient
}

// ip, password
func InitClient(params ...string) *redis.Client {
	ip := consts.RedisIpPort
	password := consts.RedisPassWord
	if len(params) >= 1 {
		ip = params[0]
	}
	if len(params) >= 2 {
		password = params[1]
	}
	return redis.NewClient(&redis.Options{
		Addr:     ip,
		Password: password,
	})
}
