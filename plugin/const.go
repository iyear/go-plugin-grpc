package plugin

type LogLevel int32

const (
	LogLevelDebug LogLevel = 0
	LogLevelInfo  LogLevel = 1
	LogLevelWarn  LogLevel = 2
	LogLevelError LogLevel = 3
)

type UnbindReason int32

const (
	UnbindExit      UnbindReason = 0 // 插件退出
	UnbindUnUsed    UnbindReason = 1 // 插件不再使用
	UnbindUpgrade   UnbindReason = 2 // 插件升级
	UnbindDowngrade UnbindReason = 3 // 插件降级
	UnbindPanic     UnbindReason = 4 // 插件异常
)

type Status int32

const (
	StatusConnected    Status = 0
	StatusDisconnected Status = 1
)
