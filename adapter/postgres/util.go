package postgres

import (
	"clean_arch/domain/model"
	"fmt"
	"time"

	pq "github.com/lib/pq"
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
		return nil, fmt.Errorf("HandleUserPqErr: %w", err)
	}
	return nil, err
}
