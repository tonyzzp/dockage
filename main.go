package main

import (
	"dockage/cmds"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	var app = &cli.App{
		Name: "dockage",
		Commands: []*cli.Command{
			cmds.CmdStart,
		},
	}
	app.Run(os.Args)
}
