package logger

import (
	"errors"
	"log"

	"go.uber.org/zap"

	"clean_arch/infra/config"
	lgr "clean_arch/infra/logger"
)

// NewLogger -
func NewLogger(c *config.Config) (lgr.LogInfoFormat, error) {
	if c.Log.Code == "zap" {
		z, er := NewZapLogger(c)
		if er != nil {
			log.Fatalf("can't initialize zap logger: %v", er)
			return nil, er
		}
		return &lggr{zapSugarLogger: z}, nil

	}
	return nil, errors.New("logger not supported : " + c.Log.Code)
}

type lggr struct {
	zapSugarLogger *zap.SugaredLogger
}

func (l *lggr) Debug(args ...interface{}) {
	l.zapSugarLogger.Debug(args)
}

func (l *lggr) Info(args ...interface{}) {
	l.zapSugarLogger.Info(args)
}

func (l *lggr) Warn(args ...interface{}) {
	l.zapSugarLogger.Warn(args)
}

func (l *lggr) Error(args ...interface{}) {
	l.zapSugarLogger.Error(args)
}

func (l *lggr) Panic(args ...interface{}) {
	l.zapSugarLogger.Panic(args)
}

func (l *lggr) Fatal(args ...interface{}) {
	l.zapSugarLogger.Fatal(args)
}
func (l *lggr) Debugf(template string, args ...interface{}) {
	l.zapSugarLogger.Debugf(template, args)
}

func (l *lggr) Infof(template string, args ...interface{}) {
	l.zapSugarLogger.Infof(template, args)
}

func (l *lggr) Warnf(template string, args ...interface{}) {
	l.zapSugarLogger.Warnf(template, args)
}

func (l *lggr) Errorf(template string, args ...interface{}) {
	l.zapSugarLogger.Errorf(template, args)
}

func (l *lggr) Panicf(template string, args ...interface{}) {
	l.zapSugarLogger.Panicf(template, args)
}

func (l *lggr) Fatalf(template string, args ...interface{}) {
	l.zapSugarLogger.Fatalf(template, args)
}
