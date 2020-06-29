package handler

import (
	"github.com/pkg/errors"

	"github.com/timoth-y/microgen/core/codegen"
	"github.com/timoth-y/microgen/core/config"
	"github.com/timoth-y/microgen/core/plugin"
)

func Generate(cfg *config.Config, option ...Option) error {
	plugins := make([]plugin.Plugin, 0)
	for _, o := range option {
		o(cfg, &plugins)
	}

	for _, p := range plugins {
		if mut, ok := p.(plugin.DirectoryGenerator); ok {
			err := mut.GenerateDirectory(); if err != nil {
				return errors.Wrap(err, p.Name())
			}
		}
	}
	// Merge again now that the generated models have been injected into the typemap
	data, err := codegen.BuildData(cfg)
	if err != nil {
		return errors.Wrap(err, "config parsing failed")
	}

	for _, p := range plugins {
		if mut, ok := p.(plugin.CodeGenerator); ok {
			err := mut.GenerateCode(data)
			if err != nil {
				return errors.Wrap(err, p.Name())
			}
		}
	}

	return nil
}
