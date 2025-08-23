package config

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm/logger"
)

type ZerologGormLogger struct {
	LogLevel logger.LogLevel
}

func NewZerologGormLogger(level logger.LogLevel) ZerologGormLogger {
	return ZerologGormLogger{
		LogLevel: level,
	}
}

func (l ZerologGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l ZerologGormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		log.Info().Msgf(msg, data...)
	}
}

func (l ZerologGormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		log.Warn().Msgf(msg, data...)
	}
}

func (l ZerologGormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		log.Error().Msgf(msg, data...)
	}
}

func (l ZerologGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	if err != nil {
		log.Error().
			Err(err).
			Str("sql", sql).
			Int64("rows", rows).
			Dur("elapsed", elapsed).
			Msg("SQL execution failed")
		return
	}

	if elapsed > 200*time.Millisecond {
		log.Warn().
			Str("sql", sql).
			Int64("rows", rows).
			Dur("elapsed", elapsed).
			Msg("Slow query detected")
	}
}
