package commands

import (
	"github.com/urfave/cli"
)

var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "c",
		Usage: "specify the configuration `file`",
	},
	cli.BoolFlag{
		Name:  "d",
		Usage: "debug",
	},
	cli.StringFlag{
		Name:  "endpoint, e",
		Usage: "Aliyun API `endpoint`",
	},
	cli.StringFlag{
		Name:  "ak",
		Usage: "Aliyun API `accessKey`",
	},
	cli.StringFlag{
		Name:  "aks",
		Usage: "Aliyun API `accessKeySecret`",
	},
}

func AddGlobalFlags(app *cli.App, flags []cli.Flag) {
	app.Flags = flags
}

func AddCmd(app *cli.App, command *cli.Command) {
	app.Commands = append(app.Commands, *command)
}
