package logger

import "go.uber.org/zap"

var instant *zap.SugaredLogger

func New() *zap.SugaredLogger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	return logger.Sugar()
}

func Instant() *zap.SugaredLogger {
	if instant == nil {
		instant = New()
	}

	return instant
}
