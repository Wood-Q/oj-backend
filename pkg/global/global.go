// pkg/global/global.go
package global

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	Db       *gorm.DB
	RedisDb  *redis.Client
	RocketMQ rocketmq.Producer
)
