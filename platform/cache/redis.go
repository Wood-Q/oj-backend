package cache

import (
	"OJ/pkg/global"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2/log"
)

func InitRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0, // use default DB
		Password: "",
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Error("Failed to connect to Redis: %v", err)
	} else {
		log.Info("Successfully connected to Redis")
	}
	global.RedisDb = RedisClient
}
