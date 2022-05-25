package codec

import (
	"fmt"
	"github.com/iyear/go-plugin-grpc/internal/pb"
	"github.com/vmihailenco/msgpack/v5"
)

func Encode(v interface{}) ([]byte, pb.CodecType, error) {
	if v == nil {
		return make([]byte, 0), pb.CodecType_Bytes, nil
	}

	switch t := v.(type) {
	case map[string]interface{}:
		bytes, err := msgpack.Marshal(t)
		if err != nil {
			return nil, 0, err
		}
		return bytes, pb.CodecType_Map, nil
	case []byte:
		return t, pb.CodecType_Bytes, nil
	default:
		return nil, 0, fmt.Errorf("unsupported type: %v", t)
	}
}
