package servergen

import (
	"log"
	"os"

	"github.com/timoth-y/microgen/core/codegen"
	"github.com/timoth-y/microgen/usecase/plugin"
)

type Plugin struct {
	filename string
}

type ServerBuild struct {
	codegen.Data

	ExecPackageName     string
	ResolverPackageName string
}

var _ plugin.CodeGenerator = &Plugin{}

var _ plugin.DirectoryGenerator = &Plugin{}

func New(filename string) plugin.Plugin {
	return &Plugin{filename}
}

func (p *Plugin) Name() string {
	return "ServerGen"
}

func (p *Plugin) GenerateCode(data *codegen.Data) error {
	log.Printf("Skipped server: %s already exists\n", p.filename)
	return nil
}

func (p *Plugin) GenerateDirectory() {
	os.Mkdir("server", os.ModePerm)
}
