package codec

import (
	"fmt"
	"github.com/iyear/go-plugin-grpc/internal/pb"
	"github.com/iyear/go-plugin-grpc/shared"
)

type Union interface {
	Map() *shared.MapConv
	Bytes() []byte
	Type() shared.CodecType
}

type nativeUnion struct {
	mmap  map[string]interface{}
	b     []byte
	ctype pb.CodecType
}

func (u *nativeUnion) Map() *shared.MapConv {
	if u.ctype != pb.CodecType_Map {
		panic("type is not map")
	}
	return shared.NewMapConv(u.mmap)
}

func (u *nativeUnion) Bytes() []byte {
	if u.ctype != pb.CodecType_Bytes {
		panic("type is not bytes")
	}
	return u.b
}

func (u *nativeUnion) Type() shared.CodecType {
	return shared.CodecType(u.ctype)
}

func (u *nativeUnion) String() string {
	switch u.ctype {
	case pb.CodecType_Map:
		return fmt.Sprintf("%v", u.mmap)
	case pb.CodecType_Bytes:
		return fmt.Sprintf("%v", u.b)
	default:
		return ""
	}
}
