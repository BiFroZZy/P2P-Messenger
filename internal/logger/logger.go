package logger

import (
	"os"

	"github.com/rs/zerolog"
)

func InitLogger() zerolog.Logger{
	log := zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()
	return log
}