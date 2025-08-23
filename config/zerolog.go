package config

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	// File untuk log info & warning
	appLogFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open app.log")
	}

	// File khusus untuk error
	errorLogFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open error.log")
	}

	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	logger := zerolog.New(consoleWriter).With().Timestamp().Logger()

	log.Logger = logger

	log.Logger = log.Logger.Hook(splitHook{
		appLogger:   zerolog.New(appLogFile).With().Timestamp().Logger(),
		errorLogger: zerolog.New(errorLogFile).With().Timestamp().Logger(),
	})
}

// splitHook untuk memisahkan log ke file
type splitHook struct {
	appLogger   zerolog.Logger
	errorLogger zerolog.Logger
}

func (h splitHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	switch {
	case level == zerolog.InfoLevel || level == zerolog.WarnLevel:
		h.appLogger.WithLevel(level).Msg(msg)
	case level >= zerolog.ErrorLevel:
		h.errorLogger.WithLevel(level).Msg(msg)
	}
}
