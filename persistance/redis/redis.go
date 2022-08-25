package redisserviceprovider

import (
	"github.com/go-redis/redis/v9"
	"github.com/themartes/erd/config"
)

type RedisInstance struct {
	Client *redis.Client
}

func (r RedisInstance) GetClient() *redis.Client {
	if r.Client == nil {
		r.Client = redis.NewClient(config.RedisConfig)
	}

	return r.Client
}
