package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/fatih/color"
	"github.com/kangaloo/cloudcli/commands/flagscheck"
	"github.com/kangaloo/cloudcli/display"
	"github.com/urfave/cli"
)

// List print a buckets list
func ListBucket(c *cli.Context) error {

	//necessary := []string{"b"}
	//conflicts := [][]string{}

	// 检查上面两个切片中的参数是否都是定义过的
	// DefinedCheck

	/*
		// 必要参数检查
		if err := cloud.NecessaryCheck(c, necessary...); err != nil {
			if err := cli.ShowSubcommandHelp(c); err != nil {
				return err
			}

			return err
		}
	*/

	client, err := NewOssClient(c.App.Metadata["config"])

	if err != nil {
		return err
	}

	list, err := client.ListBuckets()

	if err != nil {
		return err
	}

	for _, bucket := range list.Buckets {
		fmt.Println(bucket.Name)
	}

	/*
		// 冲突参数检查
		if err := cloud.ConflictCheck(c, conflicts); err != nil {
			return err
		}
	*/

	// 检查当前子命令需要的特殊参数是否满足 没有必要的
	// return specialFlagErr

	return nil
}

// ListFiles list all objects in a bucket
func ListObjects(c *cli.Context) error {
	// 必要参数检查
	// 冲突参数检查
	// 特殊参数检查

	necessary := []string{"b"}

	if err := flagscheck.NecessaryCheck(c, necessary...); err != nil {
		return err
	}

	client, err := NewOssClient(c.App.Metadata["config"])

	if err != nil {
		return err
	}

	bucket, err := client.Bucket(c.String("b"))

	if err != nil {
		return nil
	}

	objects, err := listAllObjs(bucket)
	if err != nil {
		return err
	}

	fmt.Println(color.New(color.FgHiCyan).Sprint("    size  object"))

	for _, obj := range objects {
		fmt.Printf(
			"%s  %s\n",
			color.New(color.FgHiBlack).SprintfFunc()("%s", fmt.Sprintf("%8s", "["+display.SmartSize(obj.Size)+"]")),
			obj.Key,
		)
	}

	return nil
}

func listAllObjs(bucket *oss.Bucket) ([]oss.ObjectProperties, error) {
	var objects []oss.ObjectProperties

	marker := oss.Marker("")

	for {
		objs, err := bucket.ListObjects(marker, oss.MaxKeys(1000))

		if err != nil {
			return nil, err
		}

		marker = oss.Marker(objs.NextMarker)
		objects = append(objects, objs.Objects...)

		if !objs.IsTruncated {
			break
		}
	}

	return objects, nil
}
