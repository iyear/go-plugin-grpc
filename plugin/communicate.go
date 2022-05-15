package plugin

import (
	"context"
	"fmt"
	"github.com/iyear/go-plugin-grpc/internal/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

func (p *Plugin) recv(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			msg, err := p.clients.comm.Recv()
			if err != nil {
				return fmt.Errorf("communicate recv error: %v", err)
			}

			switch msg.Type {
			case pb.CommunicateType_ExecRequest:
				go p.exec(msg.Data)
			}
		}
	}
}

func (p *Plugin) exec(data []byte) {
	req := pb.CommunicateExecRequest{}
	if err := proto.Unmarshal(data, &req); err != nil {
		p.Log.Errorf("communicate exec req unmarshal error: %v", err)
		return
	}

	result, err := p.execFunc(&req)

	resp := &pb.CommunicateExecResponse{}
	if err != nil {
		t := err.Error()
		s := &t
		resp = &pb.CommunicateExecResponse{
			ID:     req.ID,
			Result: nil,
			Err:    s,
		}
	} else {
		respb, err := structpb.NewStruct(result)
		if err != nil {
			p.Log.Errorf("communicate exec result marshal error: %v", err)
			return
		}
		resp = &pb.CommunicateExecResponse{
			ID:     req.ID,
			Result: respb,
			Err:    nil,
		}
	}

	msg, err := proto.Marshal(resp)
	if err != nil {
		p.Log.Errorf("communicate marshal exec response error: %v", err)
		return
	}

	if err = p.clients.comm.Send(&pb.CommunicateMsg{
		Type: pb.CommunicateType_ExecResponse,
		Data: msg,
	}); err != nil {
		p.Log.Errorf("communicate send exec response error: %v", err)
		return
	}
}

//exec 执行函数
func (p *Plugin) execFunc(req *pb.CommunicateExecRequest) (map[string]interface{}, error) {
	f, ok := p.handlers.Load(req.FuncName)
	if !ok {
		return nil, fmt.Errorf("func %s not found", req.FuncName)
	}
	result, err := f.(HandlerFunc)(&nativeCtx{
		plugin: p,
		args:   req.Args.AsMap(),
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
