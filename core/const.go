package core

type LogLevel int32

const (
	LogLevelDebug LogLevel = 0
	LogLevelInfo  LogLevel = 1
	LogLevelWarn  LogLevel = 2
	LogLevelError LogLevel = 3
)

type Status int32

const (
	StatusLaunched Status = 0
	StatusStopped  Status = 1
)
