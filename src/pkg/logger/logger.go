package logger

import (
	"os"
	"time"

	"github.com/charmbracelet/log"
)

var logger *log.Logger

func Get() *log.Logger {
	if logger == nil {
		return New(log.InfoLevel)
	}

	return logger
}

func New(level log.Level) *log.Logger {
	if logger != nil {
		return logger
	}

	logger = log.NewWithOptions(os.Stderr, log.Options{
		Level:           level,
		ReportCaller:    true,
		ReportTimestamp: level != log.DebugLevel,
		TimeFormat:      time.Kitchen,
		Prefix:          "Galaxy ",
	})

	return logger
}
