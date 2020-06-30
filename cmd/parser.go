package cmd

import (
	"github.com/urfave/cli/v2"

	context "github.com/timoth-y/microgen/core/context"
	"github.com/timoth-y/microgen/core/meta"
)

func NewContext(cli *cli.Context) *context.Context {
	ctx := &context.Context{}
	ctx.ApplyParams(meta.ContextParam{
		Verbose: cli.Bool("verbose"),
		Config: cli.String("config"),
	})
	return ctx
}

