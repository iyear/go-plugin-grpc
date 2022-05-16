package core

import (
	"github.com/iyear/go-plugin-grpc/conv"
	"github.com/iyear/go-plugin-grpc/internal/codec"
	"github.com/iyear/go-plugin-grpc/shared"
)

type Result interface {
	Map() *conv.MapConv
	Bytes() []byte
	Type() shared.CodecType
}
type nativeResult struct {
	*codec.Union
}
