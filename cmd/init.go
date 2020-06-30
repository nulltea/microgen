package cmd

import (
	"github.com/urfave/cli/v2"

	"github.com/timoth-y/microgen/core/context"
	"github.com/timoth-y/microgen/core/handler"
	"github.com/timoth-y/microgen/core/plugin/apigen"
	"github.com/timoth-y/microgen/core/plugin/coregen"
	"github.com/timoth-y/microgen/core/plugin/maingen"
	"github.com/timoth-y/microgen/core/plugin/servergen"
	"github.com/timoth-y/microgen/core/plugin/usecasegen"
)

var initCMD = &cli.Command{
	Name:  "init",
	Usage: "Initialize microservice project via generating base layout structure",
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "verbose, v", Usage: "show logs"},
		&cli.StringFlag{Name: "config, c", Usage: "the config filename or path"},
	},
	Action: func(ctx *cli.Context) error {
		return InitializeStructure(NewContext(ctx))
	},
}

func InitializeStructure(ctx *context.Context) error {
	options := handler.AddPlugins(coregen.NewPlugin()).
		AddPlugins(coregen.NewPlugin()).
		AddPlugins(usecasegen.NewPlugin()).
		AddPlugins(apigen.NewPlugin()).
		AddPlugins(servergen.NewPlugin("")).
		AddPlugins(maingen.NewPlugin())

	return handler.Generate(ctx, options...)
}
