package logger

import (
	"github.com/paul-ss/pgram-backend/internal/pkg/config"
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	file *os.File
}

var lg *Logger

func newLogger() *Logger {
	conf := config.C().Logger

	logrus.SetReportCaller(true)

	switch conf.Level {
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logrus.SetLevel(logrus.FatalLevel)
	default:
		logrus.SetLevel(logrus.DebugLevel)
		logrus.Error("Logger level doesn't provided")
	}

	if conf.JSON {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{})
	}

	if conf.Stdout {
		logrus.SetOutput(os.Stdout)
		return &Logger{}
	}

	file, err := os.OpenFile(conf.Filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("can't open config file: " + err.Error())
	}
	logrus.SetOutput(file)

	return &Logger{
		file: file,
	}
}

// Init can panic
func Init() func() {
	if lg != nil {
		return func() {}
	}

	lg = newLogger()
	return lg.teardown
}

func (l *Logger) teardown() {
	if l == nil {
		logrus.Error("attempt to teardown nil logger")
		return
	}

	if l.file != nil {
		logrus.SetOutput(os.Stdout)

		if err := l.file.Close(); err != nil {
			logrus.Error(err.Error())
		}

		l.file = nil
	}

	logrus.Info("logger teardown complete")
}
