package gofiber_extend

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func zapLogger(logger *zap.Logger) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		start := time.Now().Local()
		chainErr := c.Next()
		if chainErr != nil {
			logger.Error(chainErr.Error())
		}
		stop := time.Now().Local()

		fields := []zap.Field{
			zap.Int("pid", os.Getpid()),
			zap.String("elaps", stop.Sub(start).String()),
			zap.String("ip", c.IP()),
			zap.String("requestid", c.Locals("requestid").(string)),
			zap.String("userid", c.Locals("userid").(string)),
			zap.Int("status", c.Response().StatusCode()),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.String("body", string(c.Request().Body())),
		}

		if chainErr != nil {
			formatErr := chainErr.Error()
			fields = append(fields, zap.String("error", formatErr))
			logger.With(fields...).Error(formatErr)
		}
		logger.With(fields...).Info("api.request")

		return nil
	}
}
