package registry

import (
	"errors"

	"github.com/jmoiron/sqlx"

	"clean_arch/infra"
	"clean_arch/infra/config"
	"clean_arch/infra/database"
	"clean_arch/infra/logger"
	"clean_arch/infra/util"
)

var (
	Cfg *infra.Config
	Log infra.LogInfoFormat
	//Db  infra.DB
	Db *sqlx.DB
)

// InitConfig -
// params dir is used for better testing
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
	dbc, err := database.NewDBx(Cfg)
	util.FailedIf(err)
	Db = dbc
}

// InitGlobals -
func InitGlobals(dir string) {
	InitConfig(dir)
	InitLogger()
	InitDatabase()
}
