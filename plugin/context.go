package plugin

import (
	"github.com/iyear/go-plugin-grpc/conv"
	"github.com/iyear/go-plugin-grpc/internal/codec"
	"github.com/iyear/go-plugin-grpc/shared"
)

type Context interface {
	Plugin() *Plugin
	Map() *conv.MapConv
	Bytes() []byte
	Type() shared.CodecType
	L() *Logger // Logger
}

type nativeCtx struct {
	plugin *Plugin
	*codec.Union
}

func (c *nativeCtx) Plugin() *Plugin {
	return c.plugin
}

func (c *nativeCtx) L() *Logger {
	return c.plugin.Log
}
