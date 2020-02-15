package util

import (
	"strconv"
)

// String2Int64 -
func String2Int64(input string) (int64, error) {
	res, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return -1, err
	}
	return res, nil
}

// String2Int -
func String2Int(input string) (int, error) {
	res, err := strconv.Atoi(input)
	if err != nil {
		return -1, err
	}
	return res, nil
}

// Int642String -
func Int642String(input int64) string {
	return strconv.FormatInt(input, 10)
}

// Int2String -
func Int2String(input int) string {
	return strconv.Itoa(input)
}
