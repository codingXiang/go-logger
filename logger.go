package logger

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

type (
	// log 介面
	LoggerInterface interface {
		//GetLogger : 取得 Logger
		GetLogger() *logrus.Logger
		//設定 log 等級
		SetLevel(level string)
		//取得 log 等級
		GetLevel() string
		//取得 log 輸出格式
		GetFormatter() string
		//設定 log 輸出格式
		SetFormatter(format string)
		// 輸出 debug 等級 log
		Debug(args ...interface{})
		// 輸出 info 等級 log
		Info(args ...interface{})
		// 輸出 warn 等級 log
		Warn(args ...interface{})
		// 輸出 error 等級 log
		Error(args ...interface{})
		// 輸出 fatal 等級 log
		Fatal(args ...interface{})
		// 輸出 panic 等級 log
		Panic(args ...interface{})
	}
	// log 等級類別
	LogLevel string
	// log 物件
	Logger struct {
		log    *logrus.Logger
		Level  string `yaml:"level"`  //等級（有 debug、info、error、fatal 與 panic）
		Format string `yaml:"format"` //格式（有 json 或 text）
	}
)

var (
	Log LoggerInterface
)

func (level LogLevel) String() string {
	return string(level)
}

func (level LogLevel) GetLevel() logrus.Level {
	switch level.String() {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return logrus.DebugLevel
	}
}

func InterfaceToLogger(data interface{}) Logger {
	var result = Logger{}
	if jsonStr, err := json.Marshal(data); err == nil {
		json.Unmarshal(jsonStr, &result)
	}
	return result
}

func NewLogger(setting Logger) LoggerInterface {
	var l = &Logger{
		log: logrus.New(),
	}

	l.log.SetFormatter(&logrus.TextFormatter{})
	l.log.SetOutput(os.Stdout)
	l.SetLevel(setting.Level)
	l.SetFormatter(setting.Format)
	l.Info(fmt.Sprintf("log level = %s", setting.Level))
	l.Info(fmt.Sprintf("log format = %s", setting.Format))
	return l
}

func NewLoggerWithConfiger(config *viper.Viper) LoggerInterface {
	var (
		l = &Logger{
			log: logrus.New(),
		}
		level  = config.GetString("log.level")
		format = config.GetString("log.format")
		path   = config.GetString("log.path")
	)

	l.log.SetFormatter(&logrus.TextFormatter{})
	l.log.SetOutput(os.Stdout)
	l.SetLevel(level)
	l.SetFormatter(format)
	l.SetOutput(path)
	l.Info("log level =", level)
	l.Info("log format =", format)
	l.Info("log path =", path)
	return l
}

func (l *Logger) SetOutput(path string) {
	if err := os.MkdirAll(path, 0777); err != nil {
		log.Fatalf("create log folder error: %v", err)
	}
	filename := path + "/" + time.Now().Format("2006-01-02") + ".log"
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		l.GetLogger().Fatal(err)
	}
	log.SetOutput(file)
}

func (l *Logger) GetLogger() *logrus.Logger {
	return l.log
}

func (l *Logger) SetLevel(level string) {
	var logLevel = LogLevel(level)
	l.log.SetLevel(logLevel.GetLevel())
}

func (l *Logger) GetLevel() string {
	return l.Level
}

func (l *Logger) GetFormatter() string {
	return l.Format
}

func (l *Logger) SetFormatter(format string) {
	switch format {
	case "json":
		l.log.SetFormatter(&logrus.JSONFormatter{})
		break
	case "text":
		l.log.SetFormatter(&logrus.TextFormatter{})
		break
	default:
		l.log.SetFormatter(&logrus.TextFormatter{})
		break
	}
}

func (l *Logger) Debug(args ...interface{}) {
	l.log.Debug(args)
}

func (l *Logger) Info(args ...interface{}) {
	l.log.Info(args)
}

func (l *Logger) Warn(args ...interface{}) {
	l.log.Warn(args)
}

func (l *Logger) Error(args ...interface{}) {
	l.log.Error(args)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.log.Fatal(args)
}

func (l *Logger) Panic(args ...interface{}) {
	l.log.Panic(args)
}
