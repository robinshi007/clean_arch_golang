package registry

import (
	"errors"

	"clean_arch/infra"
	"clean_arch/infra/config"
	"clean_arch/infra/database"
	"clean_arch/infra/logger"
	"clean_arch/infra/util"
)

var (
	Cfg *infra.Config
	Log infra.LogInfoFormat
	Db  infra.DB
)

// InitConfig -
func InitConfig(dir string) {
	conf, err := config.NewConfig(dir)
	util.FailedIf(err)
	Cfg = conf
}

// InitLogger -
func InitLogger() {
	if Cfg == nil {
		panic(errors.New("Config is not initialized"))
	}
	logr, err := logger.NewLogger(Cfg)
	util.FailedIf(err)
	Log = logr
}

// InitDatabase -
func InitDatabase() {
	if Cfg == nil {
		panic(errors.New("Config is not initialized"))
	}
	dbc, err := database.NewDB(Cfg)
	util.FailedIf(err)
	Db = dbc
}

// InitGlobals -
func InitGlobals(dir string) {
	InitConfig(dir)
	InitLogger()
	InitDatabase()
}
