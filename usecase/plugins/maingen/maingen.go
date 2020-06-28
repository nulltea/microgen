package maingen

import (
	"github.com/timoth-y/microgen/usecase/codegen"
	"github.com/timoth-y/microgen/usecase/plugins"
)

type plugin struct {
	filename string
}

var _ plugins.CodeGenerator = &plugin{}

func NewPlugin(filename string) plugins.Plugin {
	return &plugin{filename}
}

func (p *plugin) Name() string {
	return "maingen"
}

func (p *plugin) GenerateCode(data *codegen.Data) error {
	return nil
}
