package main

import (
	logging "github.com/ipfs/go-log/v2"
	"github.com/tddey01/swifts/build"
	"github.com/tddey01/swifts/cmd"
	"github.com/urfave/cli/v2"
	"os"
)

var log = logging.Logger("advmgr")

func main() {
	log.Info("sss")
	local := []*cli.Command{
		cmd.RunCmd,
	}

	app := &cli.App{
		Name:                 "run",
		Usage:                "Remote miner worker",
		Version:              build.UserVersion(),
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "start",
				Usage: "Specify worker repo path. flag %s and env WORKER_PATH are DEPRECATION, will REMOVE SOON",
			},
		},

		After: func(c *cli.Context) error {
			if r := recover(); r != nil {
				panic(r)
			}
			return nil
		},
		Commands: local,
	}
	app.Setup()

	if err := app.Run(os.Args); err != nil {
		log.Warnf("%+v", err)
		return
	}
}
