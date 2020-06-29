package maingen

import (
	"github.com/timoth-y/microgen/core/codegen"
	"github.com/timoth-y/microgen/core/plugin"
)

type _plugin struct {
	filename string
}

var _ plugin.CodeGenerator = &_plugin{}

func NewPlugin(filename string) plugin.Plugin {
	return &_plugin{filename}
}

func (p *_plugin) Name() string {
	return "maingen"
}

func (p *_plugin) GenerateCode(data *codegen.Data) error {
	return nil
}
