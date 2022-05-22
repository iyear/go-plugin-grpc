package core

import (
	"github.com/iyear/go-plugin-grpc/internal/codec"
	"github.com/iyear/go-plugin-grpc/shared"
)

type Result interface {
	Map() *shared.MapConv   // get MapConv when result.CodecType = Map
	Bytes() []byte          // get Bytes when result.CodeType = Bytes
	Type() shared.CodecType // get CodecType
}

type nativeResult struct {
	*codec.Union
}

type Union interface {
	Map() *shared.MapConv   // get MapConv when CodecType = Map
	Bytes() []byte          // get Bytes when CodeType = Bytes
	Type() shared.CodecType // get CodecType
}
