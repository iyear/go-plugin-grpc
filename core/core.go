package core

import (
	mapset "github.com/deckarep/golang-set"
	"github.com/iyear/go-plugin-grpc/internal/pb"
	"github.com/iyear/go-plugin-grpc/internal/util"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
	"sync"
)

type Core struct {
	impl     *impl
	token    string
	plugins  sync.Map      // map[string]*PluginInfo
	opts     Options       // options
	server   *grpc.Server  // grpc server
	status   pb.CoreStatus // status
	cron     *cron.Cron    // health check
	execResp sync.Map      // exec response map : map[uint64]chan execResp
	cancel   func()        // global ctx cancel function
}

type impl struct {
	core *Core
}

type PluginInfo struct {
	name     string
	version  string
	health   int64
	shutdown chan struct{}
	impl     string // name of plugin implementation
	funcs    mapset.Set

	comm pb.Conn_CommunicateServer
}

type Interface map[string][]string

func New(token string, opts ...Option) *Core {
	// TODO ExecTimeout 应当小于等于 ServerOpts 的超时时间
	c := Core{
		token:    token,
		plugins:  sync.Map{},
		status:   pb.CoreStatus_Stopped,
		execResp: sync.Map{},
		cron:     cron.New(),
		opts:     defaultOpts(),
	}

	for _, opt := range opts {
		opt.apply(&c.opts)
	}

	c.server = grpc.NewServer(c.opts.ServerOpts...)

	i := impl{core: &c}
	c.impl = &i
	pb.RegisterConnServer(c.server, &i)
	return &c
}

func (c *Core) Token() string {
	return c.token
}

func (c *Core) Status() Status {
	return Status(c.status)
}

// Opts returns the options of the core,it's read-only
func (c *Core) Opts() Options {
	return c.opts
}

//Plugin returns the plugin info of the specified plugin
func (c *Core) Plugin(name, version string) (*PluginInfo, bool) {
	p, ok := c.plugins.Load(util.GenKey(name, version))
	if !ok {
		return nil, false
	}
	return p.(*PluginInfo), true
}

//Plugins returns the plugin info of all plugins
func (c *Core) Plugins() []*PluginInfo {
	plugins := make([]*PluginInfo, 0)
	c.plugins.Range(func(key, value interface{}) bool {
		plugins = append(plugins, value.(*PluginInfo))
		return true
	})
	return plugins
}
