package oss

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/kangaloo/cloudcli/commands/flagscheck"
	"github.com/kangaloo/cloudcli/display"
	"github.com/urfave/cli"
)

// CreateBucket create a bucket
func CreateBucket(c *cli.Context) error {

	var (
		client *oss.Client
		exist  bool
		err    error
	)

	necessary := []string{"b"}

	if err = flagscheck.NecessaryCheck(c, necessary...); err != nil {
		return err
	}

	if client, err = NewOssClient(c.App.Metadata["config"]); err != nil {
		return err
	}

	if err = client.CreateBucket(c.String("b")); err != nil {
		return err
	}

	if exist, err = client.IsBucketExist(c.String("b")); err != nil {
		return err
	}

	if exist {
		fmt.Printf("%s bucket '%s' created\n", display.HiBlack("message:"), c.String("b"))
		return nil
	}

	return errors.New("create bucket failed")
}
