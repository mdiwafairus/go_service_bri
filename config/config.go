package config

import (
	"os"
	"time"

	"github.com/rs/zerolog/log"
)

var (
	ShortTTL time.Duration
	LongTTL  time.Duration
)

func InitTTL() {
	var err error

	shortTTLStr := os.Getenv("IDEMPOTENCY_SHORT_TTL")
	longTTLStr := os.Getenv("IDEMPOTENCY_LONG_TTL")

	// default
	if shortTTLStr == "" {
		shortTTLStr = "5s"
	}
	if longTTLStr == "" {
		longTTLStr = "5m"
	}

	ShortTTL, err = time.ParseDuration(shortTTLStr)
	if err != nil {
		log.Fatal().Err(err).Msg("invalid IDEMPOTENCY_SHORT_TTL")
	}

	LongTTL, err = time.ParseDuration(longTTLStr)
	if err != nil {
		log.Fatal().Err(err).Msg("invalid IDEMPOTENCY_LONG_TTL")
	}

	log.Info().
		Dur("ShortTTL", ShortTTL).
		Dur("LongTTL", LongTTL).
		Msg("Idempotency TTL config loaded")
}
