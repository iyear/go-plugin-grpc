package core

import (
	"fmt"
	"github.com/iyear/go-plugin-grpc/internal/pb"
	"github.com/iyear/go-plugin-grpc/shared"
	"google.golang.org/protobuf/proto"
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

				bound = true
				plugin = p
				plugin.health = time.Now().Unix() // init health time
			case pb.CommunicateType_Unbind:
				if !bound {
					return fmt.Errorf("plugin not bound")
				}
				req := pb.UnbindRequest{}
				// 解析错误断开连接
				if err = proto.Unmarshal(recv.Data, &req); err != nil {
					return err
				}
				// 解绑错误直接断开连接
				return i.core.unbind(plugin.name, plugin.version, &req)
			case pb.CommunicateType_ExecResponse:
				if !bound {
					return fmt.Errorf("plugin not bound")
				}
				resp := pb.CommunicateExecResponse{}
				if err = proto.Unmarshal(recv.Data, &resp); err != nil {
					return err
				}

				// exec response hook TODO: in goroutine or here?
				go i.core.recvExecResp(&resp)
			case pb.CommunicateType_Ping:
				if !bound {
					return fmt.Errorf("plugin not bound")
				}

				// ping hook
				i.core.opts.hook.OnPluginPing(i.core, plugin)

				plugin.health = time.Now().Unix()
			case pb.CommunicateType_Log:
				if !bound {
					return fmt.Errorf("plugin not bound")
				}
				log := pb.LogInfo{}
				if err = proto.Unmarshal(recv.Data, &log); err != nil {
					return err
				}

				// plugin log hook
				i.core.opts.hook.OnPluginLog(i.core, plugin, shared.LogLevel(log.Type), log.Message)
			}
		}
	}
}
