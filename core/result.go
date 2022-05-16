package core

import (
	"github.com/iyear/go-plugin-grpc/converter"
	"google.golang.org/protobuf/types/known/structpb"
)

type Result interface {
	Map() *converter.Converter
	Bytes() []byte
}
type nativeResult struct {
	resultMap *structpb.Struct
	bytes     []byte
}

func (r *nativeResult) Map() *converter.Converter {
	return converter.New(r.resultMap)
}

func (r *nativeResult) Bytes() []byte {
	return r.bytes
}

func (r *nativeResult) String() string {
	return r.resultMap.String()
}
