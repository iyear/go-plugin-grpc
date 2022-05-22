package core

import (
	"github.com/iyear/go-plugin-grpc/internal/codec"
)

type Result interface {
	codec.Union
}

type nativeResult struct {
	codec.Union
}
