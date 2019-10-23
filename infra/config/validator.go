package config

import (
	"clean_arch/infra"

	"github.com/pkg/errors"
)

// database code. Need to map to the database code (DataStoreConfig) in the configuration yaml file.
const (
	SQLDB string = "sqldb"
)

// constant for logger code, it needs to match log code (logConfig)in configuration
const (
	LOGRUS string = "logrus"
	ZAP    string = "zap"
)

// var logLevels = [...]string{"debug", "info", "warn", "error"}

// ValidateConfig -
func ValidateConfig(config infra.Config) error {
	err := validateDataStore(config)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateLogger(config)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func validateLogger(config infra.Config) error {
	log := config.Log
	key := log.Code
	logMsg := " in validateLogger doesn't match key = "
	if ZAP != key {
		errMsg := ZAP + logMsg + key
		return errors.New(errMsg)
	}
	if ZAP != key {
		errMsg := ZAP + logMsg + key
		return errors.New(errMsg)
	}
	return nil
}

func validateDataStore(config infra.Config) error {
	dc := config.Database
	key := dc.Code
	dcMsg := " in validateDataStore doesn't match key = "
	if SQLDB != key {
		errMsg := SQLDB + dcMsg + key
		return errors.New(errMsg)
	}

	return nil
}
