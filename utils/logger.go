package utils

import "github.com/sirupsen/logrus"

var Logger *logrus.Logger

func SetupLogger() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	Logger = logger
}
