package adapter

import "go.uber.org/zap"

type Logger struct {
	logger *zap.SugaredLogger
}

func NewLogger() *Logger {
	return &Logger{
		logger: zap.NewExample().Sugar(),
	}
}

func (l Logger) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l Logger) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l Logger) Info(msg string) {
	l.logger.Info(msg)
}

func (l Logger) Error(err error) {
	l.logger.Error(err)
}

func (l Logger) Fatal(err error) {
	l.logger.Fatal(err)
}
