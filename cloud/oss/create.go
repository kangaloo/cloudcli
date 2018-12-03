package oss

import (
	"github.com/kangaloo/cloudcli/cloud"
	"github.com/urfave/cli"
)

// CreateBucket create a bucket
func CreateBucket(c *cli.Context) error  {

	necessary := []string{"b"}

	if err := cloud.NecessaryCheck(c, necessary...); err != nil {
		return err
	}

	client, err := NewOssClient(c.App.Metadata["config"])

	if err != nil {
		return err
	}

	err = client.CreateBucket(c.String("b"))

	if err != nil {
		return err
	}

	return nil
}
