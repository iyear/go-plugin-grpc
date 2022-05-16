package plugin

import (
	"github.com/iyear/go-plugin-grpc/internal/pb"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
	"sync"
)

type Logger struct {
	plugin *Plugin
}

type HandlerFunc func(ctx Context) (interface{}, error)

type Plugin struct {
	conn     *grpc.ClientConn // grpc connection
	clients  *clients         // grpc clients
	name     string           // plugin name
	token    string           // plugin token
	opts     options          // plugin options
	version  string           // plugin version
	cron     *cron.Cron       // cron for heartbeat
	status   pb.PluginStatus  // plugin status
	handlers sync.Map         // map[string]HandlerFunc // plugin handlers
	cancel   func()           // context cancel func

	Log *Logger
}

type clients struct {
	conn pb.ConnClient
	log  pb.Conn_LogClient
	comm pb.Conn_CommunicateClient
}

func New(name, ver, token string, opts ...Option) *Plugin {
	p := Plugin{
		conn:     &grpc.ClientConn{},
		name:     name,
		version:  ver,
		token:    token,
		clients:  &clients{},
		status:   pb.PluginStatus_Disconnected,
		opts:     defaultOpts(),
		cron:     cron.New(),
		handlers: sync.Map{},
	}

	for _, opt := range opts {
		opt.apply(&p.opts)
	}

	p.Log = &Logger{plugin: &p}

	return &p
}

func (p *Plugin) Handle(funcName string, handler HandlerFunc) {
	p.handlers.Store(funcName, handler)
}
