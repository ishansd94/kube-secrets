package log

import (
	"os"

	"github.com/sirupsen/logrus"
)


func init() {
	if os.Getenv("LOG_LEVEL") == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
		Debug("log.init", "logging level is set to debug")
	}
}

func Debug(funcName string, msg string) {
	logrus.Debugf("[%s] %s", funcName, msg)
}

func Info(funcName string, msg string) {
	logrus.Infof("[%s] %s \n%s", funcName, msg)
}

func Warn(funcName string, msg string) {
	logrus.Warnf("[%s] %s \n%s", funcName, msg)
}

func Error(funcName string, msg string, err error) {
	logrus.Errorf("[%s] %s \n%s", funcName, msg, err.Error())
}

func Fatal(funcName string, msg string, err error) {
	logrus.Fatalf("[%s] %s \n%s", funcName, msg, err.Error())
}
