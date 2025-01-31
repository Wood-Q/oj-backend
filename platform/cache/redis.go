package cache

import (
	"OJ/pkg/global"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2/log"
)

func InitRedis(addr,port string,db int) {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", addr, port),
		DB:       db, // use default DB
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
