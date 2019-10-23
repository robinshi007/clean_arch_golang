package infra

type (
	// LogInfo -
	LogInfo interface {
		Debug(args ...interface{})
		Info(args ...interface{})
		Warn(args ...interface{})
		Error(args ...interface{})
		Panic(args ...interface{})
		Fatal(args ...interface{})
	}

	// LogFormat -
	LogFormat interface {
		Debugf(template string, args ...interface{})
		Infof(template string, args ...interface{})
		Warnf(template string, args ...interface{})
		Errorf(template string, args ...interface{})
		Panicf(template string, args ...interface{})
		Fatalf(template string, args ...interface{})
	}

	// LogInfoFormat -
	LogInfoFormat interface {
		LogInfo
		LogFormat
	}
)
