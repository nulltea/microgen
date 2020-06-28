package codegen

import "github.com/timoth-y/microgen/config"

type Data struct {
	Config          *config.Config
	Directives      []Directive
	Objects         []Object
	Interfaces      map[string]*Interface
	Types           []Type
}