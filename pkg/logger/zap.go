package logger

import (
	"github.com/jailtonjunior94/address/configs"

	"github.com/blendle/zapdriver"
	"go.uber.org/zap"
)

func NewLogger(config *configs.Config) *zap.SugaredLogger {
	if config.DevelopmentMode {
		level := zap.NewAtomicLevelAt(zap.InfoLevel)
		loggerConfig := zap.Config{
			Level:            level,
			Development:      true,
			Encoding:         "console",
			EncoderConfig:    zapdriver.NewDevelopmentEncoderConfig(),
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stdout"},
			DisableCaller:    false,
		}
		logger, _ := loggerConfig.Build()
		return logger.Sugar()
	}

	level := zap.NewAtomicLevelAt(zap.InfoLevel)
	loggerConfig := zap.Config{
		Level:            level,
		Development:      false,
		Encoding:         "json",
		EncoderConfig:    zapdriver.NewProductionEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		DisableCaller:    false,
	}
	logger, _ := loggerConfig.Build()
	return logger.Sugar()
}
