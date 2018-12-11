package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/urfave/cli"
)

// --prefix
func Upload(c *cli.Context) error {
	// 参数检查

	var (
		client *oss.Client
		bucket *oss.Bucket
		err    error
	)

	if client, err = NewOssClient(c.App.Metadata["config"]); err != nil {
		return err
	}

	return nil
}
