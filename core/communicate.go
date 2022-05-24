package core

import (
	"fmt"
	"github.com/iyear/go-plugin-grpc/internal/pb"
	"google.golang.org/protobuf/proto"
	"strings"
	"time"
)

func (i *impl) Communicate(comm pb.Conn_CommunicateServer) error {
	bound := false
	plugin := &PluginInfo{}
	for {
		select {
		case <-plugin.shutdown:
			return fmt.Errorf("core force shutdown")
		default:
			recv, err := comm.Recv() // TODO here can shutdown plugin
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

				i.core.opts.Logger.Logf("core", LogLevelInfo, "bind plugin [%s.%s],impl [%s] interface,funcs: %v", p.name, p.version, p.impl, p.funcs.String())
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
				i.core.opts.Logger.Logf("core", LogLevelInfo, "unbind plugin %s.%s, %s:%v", plugin.name, plugin.version, pb.UnbindReason_name[int32(req.Reason)], req.Msg)
				return i.core.unbind(plugin.name, plugin.version, &req)
			case pb.CommunicateType_ExecResponse:
				if !bound {
					continue
				}
				resp := pb.CommunicateExecResponse{}
				if err = proto.Unmarshal(recv.Data, &resp); err != nil {
					return err
				}

				go i.core.recvExecResp(&resp)
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
				i.core.opts.Logger.Log(strings.Join([]string{plugin.name, plugin.version}, "."), LogLevel(log.Type), log.Message)
			}
		}
	}
}
