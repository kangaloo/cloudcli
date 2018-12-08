package oss

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/kangaloo/cloudcli/commands/flagscheck"
	"github.com/kangaloo/cloudcli/file"
	"github.com/urfave/cli"
)

// --prefix
func Download(c *cli.Context) error {
	// TODO 参数检查

	necessary := []string{"b", "o", "f"}

	if err := flagscheck.NecessaryCheck(c, necessary...); err != nil {
		return err
	}

	// TODO 注意检查 f 是否为已经存在的文件，提供覆盖选项

	if !c.IsSet("overwrite") {
		if file.FileExist(c.String("f")) {
			return errors.New(fmt.Sprintf("file %s already exist", c.String("f")))
		}
	}

	client, err := NewOssClient(c.App.Metadata["config"])

	if err != nil {
		return err
	}

	bucket, err := client.Bucket(c.String("b"))

	if err != nil {
		return err
	}

	return bucket.DownloadFile(
		c.String("o"),
		c.String("f"),
		100*1024,
		oss.Progress(&ossProgressListener{}),
	)
}
