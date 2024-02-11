package internal

import (
	"github.com/panupong-develop/go-simple/configs"
	"github.com/panupong-develop/go-simple/pkg/logging"
	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	conf := configs.GetConfig()

	if conf.App.Environment == "development" {
		return logging.NewZapLogger(logging.Development, conf.App.FileLogPath)
	}

	return logging.NewZapLogger(logging.Production, conf.App.FileLogPath)
}
