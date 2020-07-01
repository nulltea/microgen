package maingen

import (
	"github.com/timoth-y/microgen/core/codegen"
	"github.com/timoth-y/microgen/usecase/plugin"
)

type Plugin struct {

}

var _ plugin.CodeGenerator = &Plugin{}

func New() plugin.Plugin {
	return &Plugin{}
}

func (p *Plugin) Name() string {
	return "MainGen"
}

func (p *Plugin) GenerateCode(data *codegen.Data) error {
	return nil
}
