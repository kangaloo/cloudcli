package interactivemode

import (
	"github.com/kangaloo/cloudcli/interactive"
	"github.com/urfave/cli"
)

var Interactive = &cli.Command{
	Name:      "interactive",
	ShortName: "inter",
	Usage:     "interactive mode",
	Action:    interactive.Run,
}
