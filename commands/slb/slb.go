package slb

import (
	"github.com/kangaloo/cloudcli/cloud/slb"
	"github.com/urfave/cli"
)

var Slb = &cli.Command{
	Name:        "slb",
	Usage:       "aliyun SLB API tool",
	Subcommands: slbSubCmds,
}

var slbSubCmds = cli.Commands{
	*listSlb,
}

var listSlb = &cli.Command{
	Name:      "list",
	ShortName: "ls",
	Usage:     "list all slb",
	Action:    slb.List,
	Flags:     listSlbFlags,
}

var listSlbFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "a",
		Usage: "list all slb",
	},
}
