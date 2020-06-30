package apigen

import (
	"os"

	"github.com/timoth-y/microgen/core/plugin"
)

type _plugin struct {
	subPlugins []plugin.Plugin
}

var _ plugin.DirectoryGenerator = &_plugin{}

func NewPlugin() plugin.Plugin {
	return &_plugin{
		make([]plugin.Plugin, 0),
	}
}

func (p *_plugin) Name() string {
	return "ApiGen"
}

func (p *_plugin) GenerateDirectory() error {
	return os.Mkdir("api", os.ModePerm)
}
