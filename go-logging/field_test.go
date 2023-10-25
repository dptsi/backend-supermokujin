package go_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "akhmad").Info("Hello world")
	logger.WithField("username", "budi").
		WithField("name", "Akhmad Budi").
		Info("Hello world")
}

func TestFields(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "akhmad",
		"name":     "Akhmad Budi",
	}).Info("Hello world")

}
