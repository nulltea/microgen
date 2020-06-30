package context

import (
	"context"

	"github.com/timoth-y/microgen/core/config"
	"github.com/timoth-y/microgen/core/meta"
)

type Context struct {
	context.Context
	Config config.Config
	Params meta.ContextParam
	Parent *context.Context
}

func (c *Context) ApplyConfig(cfg config.Config) {
	c.Config = cfg
}

func (c *Context) ApplyParams(prm meta.ContextParam) {
	c.Params = prm
}
