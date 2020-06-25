package cmd

import (
	"github.com/urfave/cli/v2"

	"github.com/timoth-y/microgen/core/meta"
)

func parseFlags(ctx *cli.Context) meta.OperationParam {
	return meta.OperationParam{
		Verbose: ctx.Bool("verbose"),
		Config: ctx.String("config"),
	}
}
