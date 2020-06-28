package servergen

import (
	"log"
	"os"

	"github.com/timoth-y/microgen/usecase/codegen"
	"github.com/timoth-y/microgen/usecase/plugins"
)

type plugin struct {
	filename string
}

type ServerBuild struct {
	codegen.Data

	ExecPackageName     string
	ResolverPackageName string
}

var _ plugins.CodeGenerator = &plugin{}

var _ plugins.DirectoryGenerator = &plugin{}

func NewPlugin(filename string) plugins.Plugin {
	return &plugin{filename}
}

func (p *plugin) Name() string {
	return "servergen"
}

func (p *plugin) GenerateCode(data *codegen.Data) error {
	log.Printf("Skipped server: %s already exists\n", m.filename)
	return nil
}

func (p *plugin) GenerateDirectory() error {
	return os.Mkdir("server", os.ModePerm)
}
