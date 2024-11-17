package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleware(c *fiber.Ctx) error {
	logger := logrus.New()

	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "02/01/2006 15:04:05",
		ForceColors:     true,
	})

	start := time.Now()

	err := c.Next()

	logEntry := logger.WithFields(logrus.Fields{
		"method":  c.Method(),
		"path":    c.Path(),
		"status":  c.Response().StatusCode(),
		"latency": time.Since(start).String(),
	})

	if err != nil {
		logEntry.WithField("error", err.Error()).Error("Request failed")
	} else {
		if c.Response().StatusCode() >= 400 && c.Response().StatusCode() < 500 {
			logEntry.Warn("Client error")
		} else {
			logEntry.Info("Request completed!")
		}
	}

	return err
}
