package logger

import (
	"go.uber.org/zap"
	"log"
)

type (
	Logger = zap.Logger
)

func New() *Logger {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("unable to create zapLogger")
	}
	defer zapLogger.Sync() // flushes buffer, if any

	return zapLogger
}
