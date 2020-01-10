package logger

import (
	"errors"
	"log"

	"clean_arch/infra"
)

// NewLogger -
func NewLogger(c *infra.Config) (infra.LogInfoFormat, error) {
	if c.Log.Code == "zap" {
		z, er := NewZapLogger(c)
		if er != nil {
			log.Fatalf("can't initialize zap logger: %v", er)
			return nil, er
		}
		return z, nil

	}
	return nil, errors.New("logger not supported: " + c.Log.Code)
}
