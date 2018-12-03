package ecs

import (
	"fmt"
	"github.com/urfave/cli"
)

func Ecs(c *cli.Context) error {
	fmt.Println("invoke ecs function")
	if !c.IsSet("i") {
		return cli.ShowSubcommandHelp(c)
	}
	return nil
}

func ListEcs(c *cli.Context) error {
	fmt.Println("invoke ecs function")
	if !c.IsSet("i") {
		return cli.ShowSubcommandHelp(c)
	}
	return nil
}

func Info(c *cli.Context) error  {
	fmt.Println("invoke ecs > info function")

	if !c.IsSet("id") {
		return cli.ShowCommandHelp(c, c.Command.Name)
	}

	fmt.Println(c.String("id"))
	return nil
}