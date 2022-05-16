package core

import (
	"errors"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/iyear/go-plugin-grpc/internal/pb"
	"github.com/iyear/go-plugin-grpc/internal/util"
	"github.com/robfig/cron/v3"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
	"net"
	"sync"
	"time"
)

type Core struct {
	impl     *impl
	token    string
	plugins  sync.Map      // map[string]*pluginInfo
	opts     options       // options
	server   *grpc.Server  // grpc server
	status   pb.CoreStatus // status
	cron     *cron.Cron    // health check
	execResp sync.Map      // exec response map : map[uint64]chan execResp
	cancel   func()        // global ctx cancel function
}

type impl struct {
	core *Core
}

type pluginInfo struct {
	name     string
	version  string
	health   int64
	shutdown chan struct{}
	impl     string // name of plugin implementation
	funcs    mapset.Set

	comm pb.Conn_CommunicateServer
}

func New(token string, opts ...Option) *Core {
	// TODO execTimeout 应当小于等于 serverOpts 的超时时间
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

	c.server = grpc.NewServer(c.opts.serverOpts...)

	i := impl{core: &c}
	c.impl = &i
	pb.RegisterConnServer(c.server, &i)
	return &c
}

func (c *Core) Serve() error {
	defer func() {
		// if errors occur, close the server
		if c.status == pb.CoreStatus_Launched {
			return
		}
		c.Shutdown()
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.opts.port))
	if err != nil {
		return err
	}

	// TODO attention
	// health check,only internal is the half of the healthTimeout
	_, err = c.cron.AddFunc(fmt.Sprintf("@every %ds", int(c.opts.execTimeout.Seconds()/2)), c.healthCheck())
	if err != nil {
		return err
	}
	c.cron.Start()

	c.status = pb.CoreStatus_Launched

	if err = c.server.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (c *Core) ShutdownPlugin(plugin, version string) error {
	if c.status != pb.CoreStatus_Launched {
		return fmt.Errorf("core is not launched")
	}

	key := util.GenKey(plugin, version)

	p, ok := c.plugins.Load(key)
	if !ok {
		return fmt.Errorf("plugin %s not found", key)
	}

	close(p.(*pluginInfo).shutdown)
	c.plugins.Delete(key)

	return nil
}

// Shutdown core
func (c *Core) Shutdown() {
	c.cancel()
	c.cron.Stop()

	c.server.GracefulStop()

	c.status = pb.CoreStatus_Stopped
}

// Call blocks until the func is executed or timeout
//
// args can be map[string]interface{} or []byte
func (c *Core) Call(plugin, version, funcName string, args map[string]interface{}) (map[string]interface{}, error) {
	p, ok := c.plugins.Load(util.GenKey(plugin, version))
	if !ok {
		return nil, fmt.Errorf("plugin %s not found", plugin)
	}

	// reduce functions not found after transmission
	if !p.(*pluginInfo).funcs.Contains(funcName) {
		return nil, fmt.Errorf("func %s not found", funcName)
	}

	id := uint64(time.Now().UnixNano())

	// set result channel
	respCh := make(chan execResp, 0)
	c.execResp.Store(id, respCh)

	// TODO support map[string]interface{} and []byte
	//var req *anypb.Any
	//switch t := args.(type) {
	//case map[string]interface{}:
	//	reqpb, err := structpb.NewStruct(t)
	//	if err != nil {
	//		return nil, err
	//	}
	//	if req, err = anypb.New(reqpb); err != nil {
	//		return nil, err
	//	}
	//case []byte:
	//
	//default:
	//	return nil, fmt.Errorf("args type error")
	//}

	argspb, err := structpb.NewStruct(args)
	if err != nil {
		return nil, err
	}
	b, err := proto.Marshal(&pb.CommunicateExecRequest{
		ID:       id,
		FuncName: funcName,
		Args:     argspb,
	})
	// failed to marshal
	if err != nil {
		return nil, err
	}
	if err = p.(*pluginInfo).comm.Send(&pb.CommunicateMsg{Type: pb.CommunicateType_ExecRequest, Data: b}); err != nil {
		return nil, err
	}

	// exec timeout
	timer := time.NewTimer(c.opts.execTimeout)
	defer timer.Stop()
	defer func() {
		close(respCh)
		c.execResp.Delete(id)
	}()

	select {
	case <-timer.C:
		return nil, fmt.Errorf("exec %s.%s.%s timeout", plugin, version, funcName)
	case result := <-respCh:
		// TODO log info result
		if result.Err != nil {
			return nil, errors.New(*result.Err)
		}

		return result.Result.AsMap(), nil
	}
}
