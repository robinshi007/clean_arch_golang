package postgres

import (
	"fmt"
	"os"
	"time"

	pq "github.com/lib/pq"

	"clean_arch/domain/model"
	"clean_arch/infra/util"
)

// TimeNow -
func TimeNow() time.Time {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	return time.Now().In(location)
}

// HandleUserPqErr -
func HandleUserPqErr(err error) (*model.User, error) {
	if err, ok := err.(*pq.Error); ok {
		util.CW(os.Stdout, util.NRed, "\"%s\"\n", err.Error())
		return nil, fmt.Errorf("HandleUserPqErr: %w", err)
	}
	return nil, err
}

// HandleAccountPqErr -
func HandleAccountPqErr(err error) (*model.UserAccount, error) {
	if err, ok := err.(*pq.Error); ok {
		util.CW(os.Stdout, util.NRed, "\"%s\"\n", err.Error())
		return nil, fmt.Errorf("HandleAccountPqErr: %w", err)
	}
	return nil, err
}
