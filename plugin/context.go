package plugin

import (
	"github.com/iyear/go-plugin-grpc/internal/codec"
	"github.com/iyear/go-plugin-grpc/shared"
)

type Context interface {
	Plugin() *Plugin        // get self
	Map() *shared.MapConv   // get MapConv when CodecType = Map
	Bytes() []byte          // get Bytes when CodeType = Bytes
	Type() shared.CodecType // get CodecType
	L() *Logger             // Log Service
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
