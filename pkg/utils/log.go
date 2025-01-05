package utils

import (
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

// Log is the global logger instance
var (
	Log  *logrus.Logger
	once sync.Once
)

// initLogger initializes the logger
func initLogger() {
	Log = logrus.New()
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.InfoLevel)
}

// GetLogger returns the singleton logger instance
func GetLogger() *logrus.Logger {
	once.Do(initLogger)
	return Log
}
