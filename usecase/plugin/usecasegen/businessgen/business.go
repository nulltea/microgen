package businessgen

import (
	"os"

	"github.com/timoth-y/microgen/usecase/plugin"
)

type Plugin struct {

}

var _ plugin.DirectoryGenerator = &Plugin{}

func New() plugin.Plugin {
	return &Plugin{

	}
}

func (p *Plugin) Name() string {
	return "BusinessGen"
}

func (p *Plugin) GenerateDirectory() {
	os.Mkdir("usecase/business", os.ModePerm)
}
