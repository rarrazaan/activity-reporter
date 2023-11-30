package dependency

import (
	"fmt"
	"mini-socmed/internal/cons"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(config Config, logger Logger) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf(cons.RedisConnectionTemplate,
			config.RedisCache.HOST,
			config.RedisCache.PORT,
		),
	})

	logger.Infof("Successfully Load Redis Client", nil)

	return client
}
