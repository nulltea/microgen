package codegen

import "github.com/timoth-y/microgen/core/config"

type Data struct {
	Config          *config.Config
	Directives      []Directive
	Objects         []Object
	Interfaces      map[string]*Interface
	Types           []Type
}

func BuildData(cfg *config.Config) (*Data, error) {
	return nil, nil
}

