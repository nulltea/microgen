package servergen

import (
	"log"
	"os"

	"github.com/timoth-y/microgen/core/codegen"
	"github.com/timoth-y/microgen/core/plugin"
)

type _plugin struct {
	filename string
}

type ServerBuild struct {
	codegen.Data

	ExecPackageName     string
	ResolverPackageName string
}

var _ plugin.CodeGenerator = &_plugin{}

var _ plugin.DirectoryGenerator = &_plugin{}

func NewPlugin(filename string) plugin.Plugin {
	return &_plugin{filename}
}

func (p *_plugin) Name() string {
	return "servergen"
}

func (p *_plugin) GenerateCode(data *codegen.Data) error {
	log.Printf("Skipped server: %s already exists\n", m.filename)
	return nil
}

func (p *_plugin) GenerateDirectory() error {
	return os.Mkdir("server", os.ModePerm)
}
