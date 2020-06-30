package handler

import (
	"github.com/pkg/errors"

	"github.com/timoth-y/microgen/core/codegen"
	"github.com/timoth-y/microgen/core/context"
	"github.com/timoth-y/microgen/core/plugin"
)

func Generate(ctx *context.Context, option ...Option) error {
	plugins := make([]plugin.Plugin, 0)
	for _, o := range option {
		o(ctx, &plugins)
	}

	data, err := codegen.BuildData(ctx.Config)
	if err != nil {
		return errors.Wrap(err, "config parsing failed")
	}

	execute(ctx, plugins, data)

	return nil
}

func execute(ctx *context.Context, plugins []plugin.Plugin, data *codegen.Data) error {
	for _, p := range plugins {
		if dir, ok := p.(plugin.DirectoryGenerator); ok {
			err := dir.GenerateDirectory(); if err != nil {
				return errors.Wrap(err, p.Name())
			}
		}
	}

	for _, p := range plugins {
		if gen, ok := p.(plugin.CodeGenerator); ok {
			err := gen.GenerateCode(data)
			if err != nil {
				return errors.Wrap(err, p.Name())
			}
		}
	}
	return nil
}
