package usecase_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
)

var WD string

func init() {
	WD, _ = os.Getwd()
	WD = filepath.Dir(WD)
}

func TestSuite(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(AccountUcaseSuite))
	suite.Run(t, new(RedirectUcaseSuite))
	suite.Run(t, new(UserUcaseSuite))
}
