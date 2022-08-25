package config

import (
	"log"
	"strconv"

	"github.com/go-redis/redis/v9"
	"github.com/themartes/erd/env"
)

var RedisConfig *redis.Options = &redis.Options{
	Addr:     env.Params.RedisURL,
	Password: env.Params.RedisPassword,
	DB:       getRedisDB(env.Params.RedisDB),
}

// getRedisDB converts string to int because .env is fucked
func getRedisDB(envValue string) int {
	output, err := strconv.Atoi(envValue)

	if err != nil {
		log.Fatal(err)
	}

	return output
}
