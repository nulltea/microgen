package plugins

import "github.com/timoth-y/microgen/usecase/codegen"

type Plugin interface {
	Name() string
}

type CodeGenerator interface {
	GenerateCode(cfg *codegen.Data) error
}

type DirectoryGenerator interface {
	GenerateDirectory() error
}

type SubPluginCollection interface {
	SubscribeSubPlugin(plugin Plugin)
	ExecuteSubPlugin() error
}

type TypeInjector interface {
	InjectType() error
}
