package cmd

import "github.com/urfave/cli/v2"

var validCMD = &cli.Command{
	Name:  "valid",
	Usage: "Validate project by Clean Architecture dependency rule",
	Flags: []cli.Flag{
	},
	Action: func(ctx *cli.Context) error {
		return nil
	},
}
