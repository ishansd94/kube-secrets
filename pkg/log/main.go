package log

import (
	"os"

	"github.com/ishansd94/stdemuxerhook"
	"github.com/sirupsen/logrus"
)

func init() {
	if os.Getenv("LOG_LEVEL") == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
		Debug("log.init", "logging level is set to debug")
	}

	logrus.AddHook(stdemuxerhook.New(logrus.StandardLogger()))
}

func Debug(process string, msg string) {
	logrus.Debugf("[%s] %s", process, msg)
}

func Info(process string, msg string) {
	logrus.Infof("[%s] %s \n", process, msg)
}

func Warn(process string, msg string) {
	logrus.Warnf("[%s] %s \n", process, msg)
}

func Error(process string, msg string, err error) {
	logrus.Errorf("[%s] %s \n %s", process, msg, err.Error())
}

func Fatal(process string, msg string, err error) {
	logrus.Fatalf("[%s] %s \n %s", process, msg, err.Error())
}
