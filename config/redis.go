package config

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

var (
	RedisCtx = context.Background()
	RedisCli *redis.Client
)

func InitRedis() {
	RedisCli = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	// check connection
	for i := 0; i < 5; i++ {
		_, err := RedisCli.Ping(RedisCtx).Result()
		if err == nil {
			log.Info().Msg("Redis Connected")
			return
		}
		log.Warn().Msgf("Redis not ready, retrying... (%d/5)", i+1)
		time.Sleep(2 * time.Second)
	}

	log.Error().Msg("failed to connect Redis after retries")
	return
}
