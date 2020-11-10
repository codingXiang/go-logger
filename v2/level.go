package logger

import "github.com/sirupsen/logrus"

type Level int

const (
	Debug Level = iota
	Info
	Warn
	Error
	Fatal
	Panic
)

func NewLevel(l string) Level {
	switch l {
	case "debug":
		return Debug
	case "info":
		return Info
	case "warn":
		return Warn
	case "error":
		return Error
	case "fatal":
		return Fatal
	case "panic":
		return Panic
	default:
		return Debug
	}
}

func (level Level) Get() logrus.Level {
	switch level {
	case Debug:
		return logrus.DebugLevel
	case Info:
		return logrus.InfoLevel
	case Warn:
		return logrus.WarnLevel
	case Error:
		return logrus.ErrorLevel
	case Fatal:
		return logrus.FatalLevel
	case Panic:
		return logrus.PanicLevel
	default:
		return logrus.DebugLevel
	}
}

func (l Level) String() string {
	return l.Get().String()
}