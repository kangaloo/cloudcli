package oss

import (
	"fmt"
	"github.com/urfave/cli"
)

func Upload(c *cli.Context) error {
	fmt.Printf("%+v\n", c.App.Metadata["config"])

	return nil
}
