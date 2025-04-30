package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type LogrusLogger struct {
	log *logrus.Logger
}

func NewLogrusLogger() *LogrusLogger {
	log := logrus.New()

	env := os.Getenv("ENVIRONMENT")

	if env == "local" {
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
			PrettyPrint:     true,
		})
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: time.RFC3339,
		})
	}

	if env == "dev" {
		log.SetLevel(logrus.TraceLevel)
	} else {
		log.SetLevel(logrus.InfoLevel)
	}

	log.SetOutput(os.Stdout)

	return &LogrusLogger{
		log: log,
	}
}

func (l *LogrusLogger) Debug(args ...any) {
	l.log.Debug(args...)
}

func (l *LogrusLogger) Debugf(format string, args ...any) {
	l.log.Debugf(format, args...)
}

func (l *LogrusLogger) Info(args ...any) {
	l.log.Info(args...)
}

func (l *LogrusLogger) Infof(format string, args ...any) {
	l.log.Infof(format, args...)
}

func (l *LogrusLogger) Warn(args ...any) {
	l.log.Warn(args...)
}

func (l *LogrusLogger) Warnf(format string, args ...any) {
	l.log.Warnf(format, args...)
}

func (l *LogrusLogger) Error(args ...any) {
	l.log.Error(args...)
}

func (l *LogrusLogger) Errorf(format string, args ...any) {
	l.log.Errorf(format, args...)
}

func (l *LogrusLogger) Fatal(args ...any) {
	l.log.Fatal(args...)
}

func (l *LogrusLogger) Fatalf(format string, args ...any) {
	l.log.Fatalf(format, args...)
}

func (l *LogrusLogger) WithField(key string, value any) LoggerItf {
	return &LogrusEntry{
		entry: l.log.WithField(key, value),
	}
}

func (l *LogrusLogger) WithFields(fields map[string]any) LoggerItf {
	return &LogrusEntry{
		entry: l.log.WithFields(fields),
	}
}
