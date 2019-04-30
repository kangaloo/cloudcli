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
	*describeSlb,
}

var listSlb = &cli.Command{
	Name:      "list",
	ShortName: "ls",
	Usage:     "list all slb",
	Action:    slb.ListLB,
	Flags:     listSlbFlags,
}

var listSlbFlags = []cli.Flag{
	cli.BoolFlag{
		Name:  "a",
		Usage: "list all slb",
	},
}

var describeSlb = &cli.Command{
	Name:      "describe",
	ShortName: "desc",
	Usage:     "describe detail information for slb",
	Action:    slb.DescribeLB,
	Flags:     describeSlbFlags,
}

var describeSlbFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "i",
		Usage: "describe detail information for slb by `id`",
	},
}
