package plugin

import (
	"github.com/iyear/go-plugin-grpc/internal/pb"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
)

type Logger struct {
	plugin *Plugin
}

type HandlerFunc func(ctx Context) (interface{}, error)

type Plugin struct {
	conn     *grpc.ClientConn       // grpc connection
	clients  *clients               // grpc clients
	name     string                 // plugin name
	token    string                 // plugin token
	opts     Options                // plugin Options
	version  string                 // plugin version
	cron     *cron.Cron             // cron for heartbeat
	status   pb.PluginStatus        // plugin status
	handlers map[string]HandlerFunc // plugin handlers. no need to use sync.Map
	cancel   func()                 // context cancel func

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
		handlers: make(map[string]HandlerFunc),
	}

	for _, opt := range opts {
		opt.apply(&p.opts)
	}

	p.Log = &Logger{plugin: &p}

	return &p
}

func (p *Plugin) Handle(funcName string, handler HandlerFunc) {
	if _, ok := p.handlers[funcName]; ok {
		panic("plugin: func " + funcName + " already exists")
	}
	p.handlers[funcName] = handler
}

func (p *Plugin) Name() string {
	return p.name
}

func (p *Plugin) Version() string {
	return p.version
}

func (p *Plugin) Token() string {
	return p.token
}

func (p *Plugin) Status() Status {
	return Status(p.status)
}

func (p *Plugin) Funcs() []string {
	funcs := make([]string, 0)
	for f := range p.handlers {
		funcs = append(funcs, f)
	}
	return funcs
}

//Opts returns the options of the plugin,it's read-only
func (p *Plugin) Opts() Options {
	return p.opts
}
