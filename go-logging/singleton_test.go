package go_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Info("Hello world")
	logrus.Warn("Hello warning")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Hello world")
	logrus.Warn("Hello warning")
}
