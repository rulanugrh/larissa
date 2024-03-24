package pkg

import (
	"io"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

type ILogrust interface {
	StartLogger(name string, nameFunc string) *logrus.Entry
}

type Logger struct {
	*logrus.Logger
}

func NewLogger() *Logger {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetLevel(logrus.DebugLevel)

	log.SetOutput(logger.Writer())
	logger.SetOutput(io.MultiWriter(os.Stdout))

	return &Logger{logger}
}

func (l *Logger) StartLogger(name string, nameFunc string) *logrus.Entry {
	return l.WithFields(logrus.Fields{
		"name": name,
		"func": nameFunc,
	})
}
