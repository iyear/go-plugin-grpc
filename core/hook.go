package core

import (
	"fmt"
	"github.com/iyear/go-plugin-grpc/shared"
	"log"
	"strings"
)

type Hook interface {
	OnPluginMount(c *Core, p *PluginInfo)
	OnPluginUnmount(c *Core, p *PluginInfo, reason shared.UnbindReason, msg *string) // msg is optional
	OnPluginPing(c *Core, p *PluginInfo)
	OnPluginLog(c *Core, p *PluginInfo, level shared.LogLevel, msg string)
	OnExecReq(c *Core, p *PluginInfo, id uint64, funcName string, args interface{})                  // args might be []byte or map[string]interface{}
	OnExecResp(c *Core, p *PluginInfo, id uint64, t shared.CodecType, result interface{}, err error) // result might be []byte or map[string]interface{}
	// TODO Core Shutdown Plugin
}

type EmptyHook struct{}

func (e EmptyHook) OnPluginMount(_ *Core, _ *PluginInfo) {}

func (e EmptyHook) OnPluginUnmount(_ *Core, _ *PluginInfo, _ shared.UnbindReason, _ *string) {}

func (e EmptyHook) OnPluginPing(_ *Core, _ *PluginInfo) {}

func (e EmptyHook) OnPluginLog(_ *Core, _ *PluginInfo, _ shared.LogLevel, _ string) {}

func (e EmptyHook) OnExecReq(_ *Core, _ *PluginInfo, _ uint64, _ string, _ interface{}) {}

func (e EmptyHook) OnExecResp(_ *Core, _ *PluginInfo, _ uint64, _ shared.CodecType, _ interface{}, _ error) {
}

type DefaultHook struct {
	logger *log.Logger
}

func (d *DefaultHook) log(prefix string, level shared.LogLevel, v ...interface{}) {
	d.logf(prefix, level, "%s", fmt.Sprint(v...))
}

var levelMap = map[shared.LogLevel]string{
	shared.LogLevelDebug: "DEBUG",
	shared.LogLevelInfo:  "INFO",
	shared.LogLevelWarn:  "WARN",
	shared.LogLevelError: "ERROR",
}

func (d *DefaultHook) logf(prefix string, level shared.LogLevel, format string, v ...interface{}) {
	d.logger.Printf("%s [%s] %s", prefix, levelMap[level], fmt.Sprintf(format, v...))
}

func (d *DefaultHook) OnPluginMount(_ *Core, p *PluginInfo) {
	d.logf("core", shared.LogLevelInfo, "plugin [%s.%s] mounts, impl [%s], funcs: %v", p.Name(), p.Version(), p.Interface(), p.Funcs())
}

func (d *DefaultHook) OnPluginUnmount(c *Core, p *PluginInfo, reason shared.UnbindReason, msg *string) {
	d.logf("core", shared.LogLevelInfo, "plugin [%s.%s] unmounts, %s:%v", p.Name(), p.Version(), reason, msg)
}

func (d *DefaultHook) OnPluginPing(c *Core, p *PluginInfo) {
	d.logf("core", shared.LogLevelInfo, "plugin [%s.%s] ping", p.Name(), p.Version())
}

func (d *DefaultHook) OnPluginLog(c *Core, p *PluginInfo, level shared.LogLevel, msg string) {
	d.log(strings.Join([]string{p.Name(), p.Version()}, "."), level, msg)
}

func (d *DefaultHook) OnExecReq(c *Core, p *PluginInfo, id uint64, funcName string, args interface{}) {
	d.logf("core", shared.LogLevelInfo, "exec request to [%s.%s], id: %d, func: %s, args: %v", p.Name(), p.Version(), id, funcName, args)
}

func (d *DefaultHook) OnExecResp(c *Core, p *PluginInfo, id uint64, t shared.CodecType, result interface{}, err error) {
	d.logf("core", shared.LogLevelInfo, "exec response from [%s.%s], id: %d, codec: %v result: %v, err: %v", p.Name(), p.Version(), id, t, result, err)
}
