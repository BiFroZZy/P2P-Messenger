package logger

import (
	"os"
	"io"
	"fmt"
	"time"
	"strings"
	"path/filepath"

	"github.com/rs/zerolog"
)

func InitLogger() zerolog.Logger{
	log := zerolog.New(LoggerOption()).With().Timestamp().Caller().Logger()
	return log
}

func LoggerOption() io.Writer{
	return zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.RFC3339,
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%s]", i))
		},
		FormatMessage: func(i interface{}) string {
			return fmt.Sprintf("| %s |", i)
		},
		FormatCaller: func(i interface{}) string {
			return filepath.Base(fmt.Sprintf("%s", i))
		},
		PartsExclude: []string{
			zerolog.TimestampFieldName,
		},
	}
}