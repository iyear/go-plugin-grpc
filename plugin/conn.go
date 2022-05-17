package plugin

import (
	"context"
	"fmt"
	"github.com/iyear/go-plugin-grpc/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

func (p *Plugin) Mount(target string, port int) error {
	//defer func() {
	//	_ = p.Shutdown(UnbindExit, nil) // TODO 错误处理
	//}()

	ctx, cancel := context.WithCancel(context.Background())
	p.cancel = cancel

	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%d", target, port), p.opts.dialOpts...)
	if err != nil {
		return err
	}

	//fmt.Println("conn success")

	p.conn = conn
	p.clients.conn = pb.NewConnClient(conn)

	// TODO log是否开另一条连接专用还在考虑，涉及到token验证，不利于复用
	logClient, err := p.clients.conn.Log(ctx, p.opts.callOpts...)
	if err != nil {
		return err
	}
	p.clients.log = logClient
	//fmt.Println("log success")

	commClient, err := p.clients.conn.Communicate(ctx, p.opts.callOpts...)
	if err != nil {
		return err
	}
	p.clients.comm = commClient
	//fmt.Println("comm success")

	// bind
	funcs := make([]string, 0)
	p.handlers.Range(func(key, value interface{}) bool {
		funcs = append(funcs, key.(string))
		return true
	})
	b, err := proto.Marshal(&pb.BindRequest{
		Token:     p.token,
		Name:      p.name,
		Version:   p.version,
		Functions: funcs,
	})
	if err != nil {
		return err
	}

	if err = p.clients.comm.Send(&pb.CommunicateMsg{Type: pb.CommunicateType_Bind, Data: b}); err != nil {
		return err
	}

	// 开始心跳
	_, err = p.cron.AddFunc(fmt.Sprintf("@every %ds", int(p.opts.heartbeat.Seconds())), p.heartbeat())
	if err != nil {
		return err
	}
	p.cron.Start()

	// 置连接状态
	p.status = pb.PluginStatus_Connected

	return p.recv(ctx)
}

//unbind 填写参数msg将在Core打印解绑原因，如不需要传入nil
func (p *Plugin) unbind(reason UnbindReason, msg *string) error {
	b, err := proto.Marshal(&pb.UnbindRequest{
		Reason:  pb.UnbindReason(reason),
		Token:   p.token,
		Name:    p.name,
		Version: p.version,
		Msg:     msg,
	})
	if err != nil {
		return err
	}

	return p.clients.comm.Send(&pb.CommunicateMsg{Type: pb.CommunicateType_Unbind, Data: b})
}

func (p *Plugin) Shutdown(reason UnbindReason, msg *string) []error {
	p.cron.Stop()

	// 关闭连接
	errs := make([]error, 0)

	if err := p.unbind(reason, msg); err != nil {
		errs = append(errs, err)
	}
	if p.conn != nil {
		if err := p.conn.Close(); err != nil {
			errs = append(errs, err)
		}
	}

	// set status to disconnected
	p.status = pb.PluginStatus_Disconnected

	p.cancel()

	return errs
}
