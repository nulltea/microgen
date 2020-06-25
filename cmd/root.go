package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/timoth-y/microgen/usecase"
)

const Version = "v0.0.1-dev"

func Execute() {
	app := cli.NewApp()
	app.Name = "microgen"
	app.Usage = "Generate base layout of microservice project using Clean Architecture approach"
	app.Description = "Generate base layout of microservice project using Clean Architecture approach"
	app.HideVersion = true
	app.Flags = []cli.Flag{
		&cli.BoolFlag{Name: "verbose, v", Usage: "show logs"},
		&cli.StringFlag{Name: "config, c", Usage: "the config filename or path"},
	}
	app.Version = Version
	app.Before = func(context *cli.Context) error {
		if context.Bool("verbose") {
			log.SetFlags(0)
		} else {
			log.SetOutput(ioutil.Discard)
		}
		return nil
	}

	handler := NewCommander(usecase.NewOperationService())
	initCMD := handler.InitCMD()
	app.Action = initCMD.Action
	app.Commands = []*cli.Command{
		initCMD,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}