package main

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {

	// file, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	// if err != nil {
	// 	logrus.Fatal(err)
	// }

	// defer file.Close()
	logger := &lumberjack.Logger{
		Filename:   "test.log",
		MaxSize:    1,
		MaxBackups: 7,
		Compress:   false,
	}

	defer logger.Close()

	logrus.SetOutput(logger)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(false)

	for i := 0; i < 100000; i++ {
		logrus.WithFields(logrus.Fields{
			"a": 1,
			"b": 2,
		}).Error("error")
		logrus.Warning("warning")
		logrus.Info("info")
		logrus.Debug("debug")
	}
}
