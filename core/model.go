package core

import "github.com/iyear/go-plugin-grpc/internal/pb"

type execReq struct {
	*pb.CommunicateExecRequest
}

type execResp struct {
	*pb.CommunicateExecResponse
}
