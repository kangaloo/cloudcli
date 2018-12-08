package oss

import (
	"fmt"
	"github.com/urfave/cli"
)

// --prefix
func Upload(c *cli.Context) error {
	fmt.Printf("%+v\n", c.App.Metadata["config"])

	return nil
}
