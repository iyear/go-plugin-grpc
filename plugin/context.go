package plugin

import (
	"github.com/iyear/go-plugin-grpc/converter"
	"google.golang.org/protobuf/types/known/structpb"
)

type Context interface {
	Plugin() *Plugin
	Map() *converter.Converter
	Bytes() []byte
	L() *Logger // Logger
}

type nativeCtx struct {
	plugin  *Plugin
	argsMap *structpb.Struct
	bytes   []byte
}

func (c *nativeCtx) Plugin() *Plugin {
	return c.plugin
}

func (c *nativeCtx) Map() *converter.Converter {
	return converter.New(c.argsMap)
}

func (c *nativeCtx) L() *Logger {
	return c.plugin.Log
}

func (c *nativeCtx) Bytes() []byte {
	return c.bytes // TODO implement
}
