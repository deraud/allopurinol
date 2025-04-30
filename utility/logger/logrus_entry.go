package logger

import (
	"github.com/sirupsen/logrus"
)

type LogrusEntry struct {
	entry *logrus.Entry
}

func (l *LogrusEntry) Debugf(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

func (l *LogrusEntry) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

func (l *LogrusEntry) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

func (l *LogrusEntry) Info(args ...interface{}) {
	l.entry.Info(args...)
}

func (l *LogrusEntry) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

func (l *LogrusEntry) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

func (l *LogrusEntry) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

func (l *LogrusEntry) Error(args ...interface{}) {
	l.entry.Error(args...)
}

func (l *LogrusEntry) Fatalf(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

func (l *LogrusEntry) Fatal(args ...interface{}) {
	l.entry.Fatal(args...)
}

func (l *LogrusEntry) WithField(key string, value interface{}) (entry LoggerItf) {
	entry = &LogrusEntry{l.entry.WithField(key, value)}
	return
}

func (l *LogrusEntry) WithFields(args map[string]interface{}) (entry LoggerItf) {
	entry = &LogrusEntry{l.entry.WithFields(args)}
	return
}
