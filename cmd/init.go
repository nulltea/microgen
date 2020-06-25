package cmd

import (
	"github.com/urfave/cli/v2"
)

func (h Handler) NewInitCMD() *cli.Command {
	return &cli.Command{
		Name: "init",
		Usage: "Initialize microservice project via generating base layout structure",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "verbose, v", Usage: "show logs"},
			&cli.StringFlag{Name: "config, c", Usage: "the config filename or path"},
			&cli.StringFlag{Name: "server", Usage: "where to write the server stub to", Value: "server.go"},
			&cli.StringFlag{Name: "schema", Usage: "where to write the schema stub to", Value: "schema.yml"},
		},
		Action: func(ctx *cli.Context) error{
			return h.service.Init(parseFlags(ctx))
		},
	}
}
