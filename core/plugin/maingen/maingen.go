package maingen

import (
	"github.com/timoth-y/microgen/core/codegen"
	"github.com/timoth-y/microgen/core/plugin"
)

type _plugin struct {

}

var _ plugin.CodeGenerator = &_plugin{}

func NewPlugin() plugin.Plugin {
	return &_plugin{}
}

func (p *_plugin) Name() string {
	return "MainGen"
}

func (p *_plugin) GenerateCode(data *codegen.Data) error {
	return nil
}
