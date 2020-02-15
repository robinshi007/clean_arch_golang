package util_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"clean_arch/infra/util"
)

func TestString2Int64(t *testing.T) {
	var expected int64 = 23
	actual, err := util.String2Int64("23")
	util.FailedIf(err)
	assert.Equal(t, expected, actual, "they should be equal")
}
func TestString2Int(t *testing.T) {
	var expected = 45
	actual, err := util.String2Int("45")
	util.FailedIf(err)
	assert.Equal(t, expected, actual, "they should be equal")
}
func TestInt642String(t *testing.T) {
	var expected = "23"
	actual := util.Int642String(int64(23))
	assert.Equal(t, expected, actual, "they should be equal")
}
func TestInt2String(t *testing.T) {
	var expected = "405"
	actual := util.Int2String(405)
	assert.Equal(t, expected, actual, "they should be equal")
}
