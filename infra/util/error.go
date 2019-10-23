package util

// FailedIf -
func FailedIf(err error) {
	if err != nil {
		panic(err)
	}
}
