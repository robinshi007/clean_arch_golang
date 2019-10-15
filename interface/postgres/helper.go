package postgres

import (
	"clean_arch/domain/model"
	"time"

	pq "github.com/lib/pq"
	"github.com/pkg/errors"
)

// TimeNow -
func TimeNow() time.Time {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	return time.Now().In(location)
}

// HandlePqErr -
func HandlePqErr(err error) (*model.User, error) {
	if err, ok := err.(*pq.Error); ok {
		return nil, errors.Wrap(err, err.Code.Name())
	}
	return nil, err
}
