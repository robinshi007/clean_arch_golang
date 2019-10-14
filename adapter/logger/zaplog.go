package logger

import (
	"errors"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"clean_arch/infra/config"
)

// NewZapLogger -
func NewZapLogger(config *config.Config) (*zap.SugaredLogger, error) {
	var cfg zap.Config

	switch strings.ToLower(config.Mode) {
	case "dev", "development":
		cfg = zap.NewDevelopmentConfig()
	case "test":
		cfg = zap.NewDevelopmentConfig()
	case "prod", "production":
		cfg = zap.NewProductionConfig()
	default:
		return nil, errors.New("logger environment not supported")
	}

	cfg.Level = zap.NewAtomicLevelAt(getLevel(config.Log.Level))
	cfg.OutputPaths = []string{config.Log.FileName}
	log, err := cfg.Build()
	if err != nil {
		return nil, errors.New("zap logger build constructs failed")
	}
	return log.Sugar(), nil
}

func getLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn", "warning":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	}
	return zap.InfoLevel
}
