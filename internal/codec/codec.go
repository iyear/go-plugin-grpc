package codec

import (
	"fmt"
	"github.com/iyear/go-plugin-grpc/internal/pb"
	"github.com/iyear/go-plugin-grpc/shared"
)

type Union interface {
	Map() *shared.MapConv   // get MapConv when CodecType = Map, otherwise panic
	Bytes() []byte          // get Bytes when CodeType = Bytes, otherwise panic
	Type() shared.CodecType // get CodecType
	Interface() interface{} // get interface{}
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

func (u *nativeUnion) Interface() interface{} {
	switch u.Type() {
	case shared.CodecTypeMap:
		return u.Map()
	case shared.CodecTypeBytes:
		return u.Bytes()
	}
	return nil
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
