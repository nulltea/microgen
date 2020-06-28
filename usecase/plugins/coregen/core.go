package coregen

import (
	"os"

	"github.com/timoth-y/microgen/usecase/plugins"
)

type plugin struct {
	subPlugins []plugins.Plugin
}

var _ plugins.DirectoryGenerator = &plugin{}

var _ plugins.SubPluginCollection = &plugin{}

func NewPlugin() plugins.Plugin {
	return &plugin{
		make([]plugins.Plugin, 0),
	}
}

func (p *plugin) Name() string {
	return "coregen"
}

func (p *plugin) GenerateDirectory() error {
	return os.Mkdir("core", os.ModePerm)
}

func (p *plugin) SubscribeSubPlugin(plugin plugins.Plugin) {
	p.subPlugins = append(p.subPlugins, plugin)
}

func (p *plugin) ExecuteSubPlugin() error {
	panic("implement me")
}
