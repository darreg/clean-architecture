package adapter

import (
	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.SugaredLogger
}

func NewLogger() *Logger {
	l, err := zap.NewProduction(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	return &Logger{
		logger: l.Sugar(),
	}
}

func (l *Logger) EnableDebug() error {
	ld, err := zap.NewDevelopment(zap.AddCaller(), zap.AddCallerSkip(1))
	if err != nil {
		return err
	}

	l.logger = ld.Sugar()
	return nil
}

func (l *Logger) Debug(msg string, args ...any) {
	l.logger.Debugw(msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.logger.Warnw(msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.logger.Infow(msg, args...)
}

func (l *Logger) Error(err error, args ...any) {
	l.logger.Errorw(err.Error(), args...)
}

func (l *Logger) Fatal(err error, args ...any) {
	l.logger.Fatalw(err.Error(), args...)
}
