package codec

import (
	"fmt"
	"github.com/iyear/go-plugin-grpc/internal/pb"
	"github.com/vmihailenco/msgpack/v5"
)

func Decode(bytes []byte, t pb.CodecType) (Union, error) {
	if bytes == nil {
		return &nativeUnion{mmap: make(map[string]interface{}), b: make([]byte, 0), ctype: t}, nil
	}

	switch t {
	case pb.CodecType_Map:
		m := make(map[string]interface{})
		if err := msgpack.Unmarshal(bytes, &m); err != nil {
			return nil, err
		}
		return &nativeUnion{mmap: m, ctype: t}, nil
	case pb.CodecType_Bytes:
		return &nativeUnion{b: bytes, ctype: t}, nil
	default:
		return nil, fmt.Errorf("unsupported type: %v", t)
	}
}
