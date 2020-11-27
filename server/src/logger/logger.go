package logger

import (
	"github.com/rs/zerolog"
	"os"
)

// Logger ...
type Logger struct {
	zerolog.Logger
}

// NewLogger ...
func NewLogger() *Logger {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	return &Logger{Logger: logger}
}
