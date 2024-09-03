package logger

import (
	"github.com/wiraphatys/intania-shop-boilerplate/pkg/config"
	"go.uber.org/zap"
)

const (
	DEV  = "development"
	PROD = "production"
)

func NewLogger(cfg config.Config) *zap.Logger {
	return newLoggerFactory(cfg.GetServer().Env)
}

func newLoggerFactory(env string) *zap.Logger {
	switch env {
	case DEV:
		return zap.Must(zap.NewDevelopment())
	case PROD:
		return zap.Must(zap.NewProduction())
	default:
		return nil
	}
}
