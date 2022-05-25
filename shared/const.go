package shared

type CodecType int32

const (
	CodecTypeMap   = 0
	CodecTypeBytes = 1
)

type UnbindReason int32

const (
	UnbindExit      UnbindReason = 0 // 插件退出
	UnbindUnUsed    UnbindReason = 1 // 插件不再使用
	UnbindUpgrade   UnbindReason = 2 // 插件升级
	UnbindDowngrade UnbindReason = 3 // 插件降级
	UnbindPanic     UnbindReason = 4 // 插件异常
)

type LogLevel int32

const (
	LogLevelDebug LogLevel = 0
	LogLevelInfo  LogLevel = 1
	LogLevelWarn  LogLevel = 2
	LogLevelError LogLevel = 3
)
