package codec

import (
	"fmt"
	"github.com/iyear/go-plugin-grpc/internal/pb"
	"github.com/iyear/go-plugin-grpc/shared"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

func Encode(v interface{}) ([]byte, pb.CodecType, error) {
	switch t := v.(type) {
	case map[string]interface{}:
		bytes := make([]byte, 0)
		r, err := structpb.NewStruct(t)
		if err != nil {
			return nil, 0, err
		}
		if bytes, err = proto.Marshal(r); err != nil {
			return nil, 0, err
		}
		return bytes, pb.CodecType_Map, nil
	case []byte:
		return t, pb.CodecType_Bytes, nil
	default:
		return nil, 0, fmt.Errorf("unsupported type: %v", t)
	}
}

type Union struct {
	smap  *structpb.Struct
	b     []byte
	ctype pb.CodecType
}

func (u *Union) Map() *shared.MapConv {
	if u.ctype != pb.CodecType_Map {
		panic("type is not map")
	}
	return shared.NewMapConv(u.smap)
}

func (u *Union) Bytes() []byte {
	if u.ctype != pb.CodecType_Bytes {
		panic("type is not bytes")
	}
	return u.b
}

func (u *Union) Type() shared.CodecType {
	return shared.CodecType(u.ctype)
}

func (u *Union) String() string {
	switch u.ctype {
	case pb.CodecType_Map:
		return u.smap.String()
	case pb.CodecType_Bytes:
		return fmt.Sprintf("%v", u.b)
	default:
		return ""
	}
}

func Decode(bytes []byte, t pb.CodecType) (*Union, error) {
	switch t {
	case pb.CodecType_Map:
		r := structpb.Struct{}
		if err := proto.Unmarshal(bytes, &r); err != nil {
			return nil, err
		}
		return &Union{smap: &r, ctype: t}, nil
	case pb.CodecType_Bytes:
		return &Union{b: bytes, ctype: t}, nil
	default:
		return nil, fmt.Errorf("unsupported type: %v", t)
	}
}
