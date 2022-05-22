package plugin

import (
	"github.com/iyear/go-plugin-grpc/internal/codec"
)

type Context interface {
	codec.Union
	Plugin() *Plugin // get self
	L() *Logger      // Log Service
}

type nativeCtx struct {
	plugin *Plugin
	codec.Union
}

func (c *nativeCtx) Plugin() *Plugin {
	return c.plugin
}

func (c *nativeCtx) L() *Logger {
	return c.plugin.Log
}
