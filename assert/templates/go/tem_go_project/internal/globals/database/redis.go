// coding: utf-8
// @Author : lryself
// @Date : 2021/4/6 20:36
// @Software: GoLand

package database

import (
	"github.com/go-redis/redis"
	"tem_go_project/internal/globals/vipers"
)

var (
	redisClient *redis.Client
)

func InitRedisClient() (err error) {
	v := vipers.GetDatabaseViper()
	redisClient = redis.NewClient(&redis.Options{
		Addr:     v.GetString("redis.addr"),
		Password: v.GetString("redis.password"),
		DB:       v.GetInt("redis.DB"),
	})
	_, err = redisClient.Ping().Result()
	return
	// Output: PONG <nil>
}

func GetRedisManager() *redis.Client {
	return redisClient
}
