package core

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/iyear/go-plugin-grpc/shared"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

type Option interface {
	apply(*Options)
}

type Options struct {
	Port          int
	HealthTimeout time.Duration
	ExecReqChSize int
	ExecTimeout   time.Duration

	hook       Hook
	serverOpts []grpc.ServerOption
	interfaces map[string]mapset.Set // requires the plugin to implement functions that satisfy at least one interface. nil if no interfaces are required
	logLevel   shared.LogLevel       // TODO core log level
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
		serverOpts:    make([]grpc.ServerOption, 0),
		HealthTimeout: time.Second * 15,
		logLevel:      shared.LogLevelInfo,
		Port:          13000,
		ExecReqChSize: 0,
		hook: &DefaultHook{
			logger: log.New(os.Stdout, "", log.LstdFlags),
		},
		interfaces:  nil, // default: no interfaces required
		ExecTimeout: time.Second * 10,
	}
}

// WithHealthTimeout set the HealthTimeout, default is 15s
//
// plugins with intervals higher than this limit will be disconnected and removed
func WithHealthTimeout(dur time.Duration) Option {
	return newOption(func(options *Options) {
		options.HealthTimeout = dur
	})
}

//WithLogLevel set the log level, default is Info
func WithLogLevel(level shared.LogLevel) Option {
	return newOption(func(options *Options) {
		options.logLevel = level
	})
}

//WithPort set the port, default is 13000
func WithPort(port int) Option {
	return newOption(func(options *Options) {
		options.Port = port
	})
}

//WithExecReqChSize set the chan size, default is 0
func WithExecReqChSize(size int) Option {
	return newOption(func(options *Options) {
		options.ExecReqChSize = size
	})
}

//WithExecTimeout set the exec timeout, default is 10s
func WithExecTimeout(dur time.Duration) Option {
	return newOption(func(options *Options) {
		options.ExecTimeout = dur
	})
}

func WithServerOpts(opts ...grpc.ServerOption) Option {
	return newOption(func(options *Options) {
		options.serverOpts = opts
	})
}

func WithInterfaces(interfaces ...Interface) Option {
	return newOption(func(options *Options) {
		intfs := make(map[string]mapset.Set)
		for _, mv := range interfaces {
			for k, v := range mv {
				intf := mapset.NewSet()
				for _, s := range v {
					intf.Add(s)
				}
				intfs[k] = intf
			}
		}
		options.interfaces = intfs
	})
}

func WithHooks(h Hook) Option {
	return newOption(func(options *Options) {
		options.hook = h
	})
}
