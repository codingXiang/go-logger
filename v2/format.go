package logger

import "github.com/sirupsen/logrus"

type Format int

const (
	Text Format = iota
	Json
)

func NewFormat(f string) Format {
	switch f {
	case "json":
		return Json
	case "text":
		return Text
	default:
		return Text
	}
}

func (f Format) Get() logrus.Formatter {
	switch f {
	case Json:
		return &logrus.JSONFormatter{}
	case Text:
		return &logrus.TextFormatter{}
	default:
		return &logrus.JSONFormatter{}
	}
}
