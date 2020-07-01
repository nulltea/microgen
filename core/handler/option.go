package handler

import (
	"github.com/timoth-y/microgen/core/context"
	"github.com/timoth-y/microgen/usecase/plugin"
)

type Option func(ctx *context.Context, plugins *[]plugin.Plugin)

type Options []Option

func NoPlugins() Option {
	return func(ctx *context.Context, plugins *[]plugin.Plugin) {
		*plugins = nil
	}
}

func AddPlugin(p plugin.Plugin) Option {
	return func(ctx *context.Context, plugins *[]plugin.Plugin) {
		*plugins = append(*plugins, p)
	}
}

func AddPlugins(p plugin.Plugin) Options {
	options := make([]Option, 0)
	options = append(options, AddPlugin(p))
	return options
}

func (o Options) AddPlugin(p plugin.Plugin) Options {
	return append(o, AddPlugin(p))
}
