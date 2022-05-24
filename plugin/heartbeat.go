package plugin

import (
	"encoding/binary"
	"github.com/iyear/go-plugin-grpc/internal/pb"
	"time"
)

func (p *Plugin) heartbeat() func() {
	return func() {
		buf := make([]byte, 4)
		binary.BigEndian.PutUint32(buf, uint32(time.Now().Unix()))

		if err := p.clients.comm.Send(&pb.CommunicateMsg{
			Type: pb.CommunicateType_Ping,
			Data: buf,
		}); err != nil {
			p.Log.Errorf("Heartbeat error: %v", err)
		}
	}
}
