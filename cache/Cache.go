package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"senducode/conf"
	"senducode/tolog"
	"strconv"
)

// 'RedisClient' is a global variable representing the Redis client.
var RedisClient *redis.Client

// 'InitCache' initializes the Redis cache connection.
func InitCache() {
	// Convert the cache DB string to an integer.
	db, _ := strconv.Atoi(conf.CacheConf.DB)

	// Create a new Redis client using the configuration from the 'conf' package.
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.CacheConf.Host, conf.CacheConf.Port),
		Password: conf.CacheConf.Password,
		DB:       db,
	})

	// Set the global 'RedisClient' variable to the created client.
	RedisClient = client

	// Ping the Redis server to check the connection.
	pong, err := client.Ping().Result()
	if err != nil {
		// Log an error if there is an issue with the connection.
		tolog.Log().Errorf("%e", err).PrintAndWriteSafe()
	}

	// Log a message indicating a successful connection to Redis.
	tolog.Log().Infof("Connected to Redis: %s", pong).PrintLog()
}
