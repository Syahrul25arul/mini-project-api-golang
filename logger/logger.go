package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	config.EncoderConfig = encoderConfig

	var err error
	// log, err = zap.NewProduction(zap.AddCallerSkip(1))
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, field ...zapcore.Field) {
	log.Info(message, field...)
}

func Error(message string, field ...zapcore.Field) {
	log.Error(message, field...)
}

func Debug(message string, field ...zapcore.Field) {
	log.Debug(message, field...)
}

func Warning(message string, field ...zapcore.Field) {
	log.Warn(message, field...)
}

func Fatal(message string, field ...zapcore.Field) {
	log.Fatal(message, field...)
}
