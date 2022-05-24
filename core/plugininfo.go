package core

//Name of plugin
func (p *PluginInfo) Name() string {
	return p.name
}

//Version of plugin
func (p *PluginInfo) Version() string {
	return p.version
}

//Interface implemented by plugin
func (p *PluginInfo) Interface() string {
	return p.impl
}

func (p *PluginInfo) Funcs() []string {
	funcs := make([]string, 0)
	for f := range p.funcs.Iter() {
		funcs = append(funcs, f.(string))
	}
	return funcs
}
