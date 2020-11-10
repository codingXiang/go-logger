package logger

import (
	"fmt"
	"github.com/codingXiang/configer/v2"
	"github.com/sirupsen/logrus"
	"github.com/snowzach/rotatefilehook"
	"github.com/spf13/viper"
	"log"
	"os"
	"runtime"
	"strings"
)

const (
	LOG      = "log"
	LEVEL    = "level"
	MAX_AGE  = "maxAge"
	FORMAT   = "format"
	FILENAME = "filename"
	PATH     = "path"
)

type Logger struct {
	config *viper.Viper
	*logrus.Logger
}

func New(config *viper.Viper) *Logger {
	var (
		l = new(Logger)
	)
	l.config = config

	l.Logger = newLogger(config)
	l.Output(config)
	l.Info(fmt.Sprintf("log level = %s", l.GetLevel().String()))
	return l
}

func Default() *Logger {
	c := configer.NewCoreWithData(defaultConfig)
	c.SetConfigType(configer.YAML.String())
	if config, err := c.ReadConfig();err == nil {
		return New(config)
	} else {
		panic(err)
	}
}

func newLogger(config *viper.Viper) *logrus.Logger {
	var (
		logger = logrus.New()
		level  = config.GetString(GetConfigPath(LOG, LEVEL))
		format = config.GetString(GetConfigPath(LOG, FORMAT))
	)

	logger.SetFormatter(&logrus.TextFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(NewLevel(level).Get())
	logger.SetFormatter(NewFormat(format).Get())
	return logger
}

func (l *Logger) Output(config *viper.Viper) {
	var (
		path   = config.GetString(GetConfigPath(LOG, PATH))
		name   = config.GetString(GetConfigPath(LOG, FILENAME))
		maxAge = config.GetInt(GetConfigPath(LOG, MAX_AGE))
	)
	if err := os.MkdirAll(path, 0777); err != nil {
		log.Fatalf("create log folder error: %v", err)
	}
	filename := path + GetPathSymbol() + name

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   filename,
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     maxAge, //days
		Level:      l.GetLevel(),
		Formatter:  l.Formatter,
	})
	if err != nil {
		log.Fatalf("open log file error: %v", err)
	}
	l.SetOutput(os.Stdout)
	l.AddHook(rotateFileHook)
}

func GetConfigPath(key string, path ...string) string {
	return key + "." + strings.Join(path, ".")
}

func GetPathSymbol() string {
	switch runtime.GOOS {
	case "windows":
		return "\\"
	default:
		return "/"
	}
}
