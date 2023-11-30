package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"senducode/conf"
	"senducode/tolog"
	"strconv"
)

var RedisClient *redis.Client

func InitCache() {
	db, _ := strconv.Atoi(conf.CacheConf.DB)
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.CacheConf.Host, conf.CacheConf.Port),
		Password: conf.CacheConf.Password,
		DB:       db,
	})

	RedisClient = client

	pong, err := client.Ping().Result()
	if err != nil {
		tolog.Log().Errorf("%e", err).PrintAndWriteSafe()
	}
	tolog.Log().Infof("Connected to Redis:%s", pong).PrintLog()

}
