package core

import (
	"errors"
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/iyear/go-plugin-grpc/internal/pb"
	"github.com/iyear/go-plugin-grpc/internal/util"
	"google.golang.org/protobuf/proto"
	"strings"
	"time"
)

func (i *impl) Communicate(comm pb.Conn_CommunicateServer) error {
	bound := false
	plugin := &pluginInfo{}
	for {
		select {
		case <-plugin.shutdown:
			return fmt.Errorf("core force shutdown")
		default:
			recv, err := comm.Recv()
			if err != nil {
				return err
			}
			switch recv.Type {
			case pb.CommunicateType_Bind:
				if bound {
					continue
				}
				req := pb.BindRequest{}
				if err = proto.Unmarshal(recv.Data, &req); err != nil {
					return err
				}
				p, err := i.core.bind(&req, comm)
				if err != nil {
					return err
				}

				i.core.opts.logger.Logf("core", LogLevelInfo, "bind plugin [%s.%s],impl [%s] interface,funcs: %v", p.name, p.version, p.impl, p.funcs.String())
				bound = true
				plugin = p
				plugin.health = time.Now().Unix() // init health time
			case pb.CommunicateType_Unbind:
				if !bound {
					continue
				}
				req := pb.UnbindRequest{}
				// 解析错误断开连接
				if err = proto.Unmarshal(recv.Data, &req); err != nil {
					return err
				}
				// 解绑错误直接断开连接
				i.core.opts.logger.Logf("core", LogLevelInfo, "unbind plugin %s.%s, %s:%v", req.Name, req.Version, pb.UnbindReason_name[int32(req.Reason)], req.Msg)
				return i.core.unbind(&req)
			case pb.CommunicateType_ExecResponse:
				if !bound {
					continue
				}
				resp := pb.CommunicateExecResponse{}
				if err = proto.Unmarshal(recv.Data, &resp); err != nil {
					return err
				}

				if err = i.core.recvExecResp(&resp); err != nil {
					return err
				}
			case pb.CommunicateType_Ping:
				if !bound {
					continue
				}
				plugin.health = time.Now().Unix()
			case pb.CommunicateType_Log:
				if !bound {
					continue
				}
				log := pb.LogInfo{}
				if err = proto.Unmarshal(recv.Data, &log); err != nil {
					return err
				}
				i.core.opts.logger.Log(strings.Join([]string{plugin.name, plugin.version}, "."), LogLevel(log.Type), log.Message)
			}
		}
	}
}

func (c *Core) bind(req *pb.BindRequest, comm pb.Conn_CommunicateServer) (*pluginInfo, error) {
	// invalid token, disconnect
	if req.Token != c.token {
		return nil, errors.New("invalid token")
	}

	// must impl only one of the interfaces
	funcs := mapset.NewSet()
	for _, f := range req.Functions {
		funcs.Add(f)
	}
	implName := ""
	if c.opts.interfaces != nil {
		impls := 0
		for name, set := range c.opts.interfaces {
			if funcs.IsSuperset(set) {
				impls++
				implName = name
			}
		}
		if impls != 1 {
			return nil, fmt.Errorf("must implement only one of the interfaces")
		}
	}

	key := util.GenKey(req.Name, req.Version)
	if _, ok := c.plugins.Load(key); ok {
		// 已存在插件断开连接
		return nil, fmt.Errorf("plugin %s.%s is exists", req.Name, req.Version)
	}

	info := pluginInfo{
		name:     req.Name,
		version:  req.Version,
		health:   0,
		shutdown: make(chan struct{}, 0),
		comm:     comm,
		impl:     implName,
		funcs:    funcs,
	}
	c.plugins.Store(key, &info)
	return &info, nil
}

func (c *Core) unbind(req *pb.UnbindRequest) error {
	if c.token != req.Token {
		return errors.New("invalid token")
	}
	key := util.GenKey(req.Name, req.Version)
	if _, ok := c.plugins.Load(key); !ok {
		return fmt.Errorf("plugin %s.%s is not exists", req.Name, req.Version)
	}
	c.plugins.Delete(key)
	return nil
}

func (c *Core) recvExecResp(resp *pb.CommunicateExecResponse) error {
	exec, ok := c.execResp.Load(resp.ID)
	if !ok {
		return fmt.Errorf("exec response channel not found: %d", resp.ID)
	}

	go func() {
		timer := time.NewTimer(time.Second * 5)
		defer timer.Stop()

		r := execResp{
			CommunicateExecResponse: &pb.CommunicateExecResponse{
				ID:     resp.ID,
				Result: resp.Result,
				Err:    resp.Err,
			}}

		select {
		case <-timer.C:
		case exec.(chan execResp) <- r:
		}
	}()

	return nil
}
