package logger

var Log LoggerItf

type LoggerItf interface {
	Debug(args ...any)
	Debugf(format string, args ...any)
	Info(args ...any)
	Infof(format string, args ...any)
	Warn(args ...any)
	Warnf(format string, args ...any)
	Error(args ...any)
	Errorf(format string, args ...any)
	Fatal(args ...any)
	Fatalf(format string, args ...any)
	WithField(key string, value any) LoggerItf
	WithFields(fields map[string]any) LoggerItf
}

func SetLogger(log LoggerItf) {
	Log = log
}
