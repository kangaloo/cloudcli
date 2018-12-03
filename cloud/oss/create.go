package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/kangaloo/cloudcli/cloud"
	"github.com/urfave/cli"
)

// CreateBucket create a bucket
func CreateBucket(c *cli.Context) error {

	var (
		client *oss.Client
		err    error
	)

	necessary := []string{"b"}

	if err = cloud.NecessaryCheck(c, necessary...); err != nil {
		return err
	}

	if client, err = NewOssClient(c.App.Metadata["config"]); err != nil {
		return err
	}

	if err = client.CreateBucket(c.String("b")); err != nil {
		return err
	}

	return nil
}
