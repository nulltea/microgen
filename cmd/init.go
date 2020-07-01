package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/timoth-y/microgen/core/context"
	"github.com/timoth-y/microgen/core/handler"
	"github.com/timoth-y/microgen/usecase/plugin/apigen"
	"github.com/timoth-y/microgen/usecase/plugin/apigen/restgen"
	"github.com/timoth-y/microgen/usecase/plugin/coregen"
	"github.com/timoth-y/microgen/usecase/plugin/coregen/modelgen"
	"github.com/timoth-y/microgen/usecase/plugin/coregen/repogen"
	"github.com/timoth-y/microgen/usecase/plugin/coregen/servicegen"
	"github.com/timoth-y/microgen/usecase/plugin/maingen"
	"github.com/timoth-y/microgen/usecase/plugin/servergen"
	"github.com/timoth-y/microgen/usecase/plugin/usecasegen"
	"github.com/timoth-y/microgen/usecase/plugin/usecasegen/businessgen"
	"github.com/timoth-y/microgen/usecase/plugin/usecasegen/storagegen"
)

var initCMD = &cli.Command{
	Name:  "init",
	Usage: "Initialize microservice project via generating base layout structure",
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "verbose, v", Usage: "show logs"},
	},
	Action: func(ctx *cli.Context) error {
		return InitializeStructure(NewContext(ctx))
	},
}

func InitializeStructure(ctx *context.Context) error {
	options := handler.AddPlugins(coregen.New()).
			AddPlugin(modelgen.New()).
			AddPlugin(repogen.New()).
			AddPlugin(servicegen.New()).
		AddPlugin(usecasegen.New()).
			AddPlugin(businessgen.New()).
			AddPlugin(storagegen.New()).
		AddPlugin(apigen.New()).
			AddPlugin(restgen.New()).
		AddPlugin(servergen.New("")).
		AddPlugin(maingen.New())

	if err := handler.Generate(ctx, options...); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		return err
	}
	fmt.Fprint(os.Stdout, "Microservice project initialized successfully")
	return nil
}
