package core

import (
	"github.com/iyear/go-plugin-grpc/internal/codec"
	"github.com/iyear/go-plugin-grpc/shared"
)

type Result interface {
	Map() *shared.MapConv
	Bytes() []byte
	Type() shared.CodecType
}
type nativeResult struct {
	*codec.Union
}
