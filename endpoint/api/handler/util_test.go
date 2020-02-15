package handler_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
)

var WD string

func init() {
	WD, _ = os.Getwd()
	WD = filepath.Dir(filepath.Dir(filepath.Dir(WD)))
}

func TestSuite(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(UserHandlerSuite))
	suite.Run(t, new(GraphQLHandlerSuite))
}
