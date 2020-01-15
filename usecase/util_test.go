package usecase_test

import (
	"os"
	"path/filepath"
)

var WD string

func init() {
	WD, _ = os.Getwd()
	WD = filepath.Dir(WD)
}
