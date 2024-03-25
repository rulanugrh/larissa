package pkg

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/rulanugrh/larissa/internal/config"
	"github.com/sirupsen/logrus"
	"github.com/yukitsune/lokirus"
)

type ILogrust interface {
	StartLogger(name string, nameFunc string) *logrus.Entry
}

type Logger struct {
	*logrus.Logger
}

func NewLogger(conf *config.App) *Logger {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetLevel(logrus.DebugLevel)

	log.SetOutput(logger.Writer())
	logger.SetOutput(io.MultiWriter(os.Stdout))

	opts := lokirus.NewLokiHookOptions().
		WithLevelMap(lokirus.LevelMap{logrus.PanicLevel: "critical"}).
		WithFormatter(&logrus.JSONFormatter{}).
		WithStaticLabels(lokirus.Labels{
			"app": "larissa",
			"env": "dev",
		})

	hook := lokirus.NewLokiHookWithOpts(
		fmt.Sprintf("http://%s:%s", conf.Loki.Host, conf.Loki.Port),
		opts,
		logrus.InfoLevel,
		logrus.ErrorLevel,
		logrus.DebugLevel,
	)

	logger.AddHook(hook)
	return &Logger{logger}
}

func (l *Logger) StartLogger(name string, nameFunc string) *logrus.Entry {
	return l.WithFields(logrus.Fields{
		"name": name,
		"func": nameFunc,
	})
}
