package cmd

import "github.com/urfave/cli/v2"

var RunCmd = &cli.Command{
	Name:  "start",
	Usage: "Start Swifts  AWS Manages",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "listen",
			Usage:   "host address and port the worker api will listen on",
			Value:   "0.0.0.0:3456",
			EnvVars: []string{"LOTUS_WORKER_LISTEN"},
		},
	},
	Action: func(cctx *cli.Context) error {
		return nil
	},
}
