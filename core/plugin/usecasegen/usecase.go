package usecasegen

import (
	"os"

	"github.com/timoth-y/microgen/core/plugin"
)

type _plugin struct {
	subPlugins []plugin.Plugin
}

var _ plugin.DirectoryGenerator = &_plugin{}

var _ plugin.SubPluginCollection = &_plugin{}

func NewPlugin() plugin.Plugin {
	return &_plugin{
		make([]plugin.Plugin, 0),
	}
}

func (p *_plugin) Name() string {
	return "usecasegen"
}

func (p *_plugin) GenerateDirectory() error {
	return os.Mkdir("usecase", os.ModePerm)
}

func (p *_plugin) SubscribeSubPlugin(plugin plugin.Plugin) {
	p.subPlugins = append(p.subPlugins, plugin)
}

func (p *_plugin) ExecuteSubPlugin() error {
	panic("implement me")
}
