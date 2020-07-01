package cmd

import "github.com/urfave/cli/v2"

var genCMD = &cli.Command{
	Name:  "gen",
	Usage: "Generate microservice project base logic",
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "verbose, v", Usage: "show logs"},
		&cli.StringFlag{Name: "config, c", Usage: "the config filename or path"},
	},
	Action: func(ctx *cli.Context) error {
		return nil
	},
}
