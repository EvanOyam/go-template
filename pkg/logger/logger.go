package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func InitLog(logPath string) (err error) {
	if file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
		logrus.SetOutput(io.MultiWriter(os.Stdout, file))
	} else {
		logrus.SetOutput(os.Stdout)
		logrus.Error(err)
		return err
	}
	return
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}

func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

func Error(args ...interface{}) {
	logrus.Error(args...)
}

func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}
