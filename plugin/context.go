package plugin

type Context interface {
	Plugin() *Plugin
	Args() map[string]interface{}
	L() *Logger // Logger
}

type nativeCtx struct {
	plugin *Plugin
	args   map[string]interface{}
}

type converter map[string]interface{}

func (c converter) Int(key string) int {
	return int(c[key].(float64))
}

func (c *nativeCtx) Plugin() *Plugin {
	return c.plugin
}

func (c *nativeCtx) Args() map[string]interface{} {
	return c.args
}

func (c *nativeCtx) L() *Logger {
	return c.plugin.Log
}
