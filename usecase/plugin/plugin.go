package plugin

import "github.com/timoth-y/microgen/core/codegen"

type Plugin interface {
	Name() string
}

type CodeGenerator interface {
	GenerateCode(cfg *codegen.Data) error
}

type DirectoryGenerator interface {
	GenerateDirectory()
}

type TypeInjector interface {
	InjectType() error
}
