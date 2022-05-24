package plugin

import (
	"google.golang.org/grpc"
	"time"
)

type Option interface {
	apply(*Options)
}

type Options struct {
	DialOpts []grpc.DialOption
	CallOpts []grpc.CallOption

	Heartbeat time.Duration
	LogLevel  LogLevel
	OnPanic   func(plugin *Plugin, execID uint64, funcName string, err error)
}

type option struct {
	f func(*Options)
}

func (o *option) apply(do *Options) {
	o.f(do)
}

func newOption(f func(options *Options)) *option { return &option{f: f} }

func defaultOpts() Options {
	return Options{
		DialOpts:  make([]grpc.DialOption, 0),
		CallOpts:  make([]grpc.CallOption, 0),
		Heartbeat: time.Second * 10,
		LogLevel:  LogLevelInfo,
		OnPanic:   func(plugin *Plugin, execID uint64, funcName string, err error) {},
	}
}

//WithHeartbeat set heartbeat, default is 10s
func WithHeartbeat(dur time.Duration) Option {
	return newOption(func(options *Options) {
		options.Heartbeat = dur
	})
}

//WithLogLevel set log level, default is Info
func WithLogLevel(level LogLevel) Option {
	return newOption(func(options *Options) {
		options.LogLevel = level
	})
}

func WithDialOpts(opts ...grpc.DialOption) Option {
	return newOption(func(options *Options) {
		options.DialOpts = opts
	})
}

func WithCallOpts(opts ...grpc.CallOption) Option {
	return newOption(func(options *Options) {
		options.CallOpts = opts
	})
}

//WithOnPanic default is Log.Errorf
func WithOnPanic(f func(plugin *Plugin, execID uint64, funcName string, err error)) Option {
	return newOption(func(options *Options) {
		options.OnPanic = f
	})
}
