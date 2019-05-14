package ecs

import (
	"github.com/kangaloo/cloudcli/cloud/ecs"
	"github.com/urfave/cli"
)

var Ecs = &cli.Command{
	Name:        "ecs",
	Usage:       "aliyun ECS API tool",
	Subcommands: ecsSubCmds,
}

var ecsSubCmds = cli.Commands{
	*listEcs,
	*describeEcs,
}

var listEcs = &cli.Command{
	Name:      "list",
	ShortName: "ls",
	Usage:     "",
	Action:    ecs.ListEcs,
	Flags:     listEcsFlags,
}

var listEcsFlags = []cli.Flag{
	cli.BoolFlag{},
}

var describeEcs = &cli.Command{
	Name:      "describe",
	ShortName: "desc",
	Usage:     "",
	Action:    ecs.DescribeEcs,
	Flags:     descEcsFlags,
}

var descEcsFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "i",
		Usage: "describe ecs by `id`",
	},
}
