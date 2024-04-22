package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

// SetUpLogger initializes the logger with JSON formatter
func SetUpLogger() {
	// Initialize logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
        ForceColors: true, // Force color output even if not writing to a terminal
		FullTimestamp: true,
    })
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)

	Logger = logger
}