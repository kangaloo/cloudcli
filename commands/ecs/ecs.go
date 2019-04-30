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
