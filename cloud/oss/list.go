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

	return nil
}

// ListFiles list all objects in a bucket
func ListObjects(c *cli.Context) error {
	// 必要参数检查
	// 冲突参数检查
	// 特殊参数检查
	// finished -b -q
	// finished -n --prefix  --all

	var (
		client  *oss.Client
		bucket  *oss.Bucket
		objects []oss.ObjectProperties
		err     error
	)

	conflict := [][]string{
		{"n", "all"},
	}

	necessary := []string{"b"}

	if err = flagscheck.AtMostOneCheck(c, conflict); err != nil {
		return err
	}

	if err = flagscheck.NecessaryCheck(c, necessary...); err != nil {
		return err
	}

	if client, err = NewOssClient(c.App.Metadata["config"]); err != nil {
		return err
	}

	if bucket, err = client.Bucket(c.String("b")); err != nil {
		return err
	}

	if objects, err = AllObjs(bucket, c); err != nil {
		return err
	}

	printObjects(c, objects)

	return nil
}

func AllObjs(bucket *oss.Bucket, c *cli.Context) ([]oss.ObjectProperties, error) {
	// -n --all
	return NumObjs(bucket, c)
}

func NumObjs(bucket *oss.Bucket, c *cli.Context) ([]oss.ObjectProperties, error) {

	var (
		objList oss.ListObjectsResult
		objs    []oss.ObjectProperties
		err     error
		marker  = oss.Marker("")
		part    = 1000
	)

	// 指定了-n参数，或者没有指定--all参数时
	if c.IsSet("n") || !c.IsSet("all") {
		if c.Int("n") <= 1000 {
			objs, err := bucket.ListObjects(
				oss.Prefix(c.String("prefix")),
				oss.MaxKeys(c.Int("n")),
			)

			return objs.Objects, err
		}

		num := c.Int("n")
		num -= part

		for {
			if objList, err = numObjs(bucket, c, marker, part); err != nil {
				return nil, err
			}

			objs = append(objs, objList.Objects...)
			marker = oss.Marker(objList.NextMarker)

			if !objList.IsTruncated {
				return objs, nil
			}

			if num == 0 {
				return objs, nil
			}

			if num <= part {
				part = num
				num = 0
			}

			if num > part {
				num -= part
			}
		}
	}

	// 指定--all参数时
	for {

		objList, err = bucket.ListObjects(
			marker,
			oss.MaxKeys(1000),
			oss.Prefix(c.String("prefix")),
		)

		if err != nil {
			return nil, err
		}

		marker = oss.Marker(objList.NextMarker)
		objs = append(objs, objList.Objects...)

		if !objList.IsTruncated {
			break
		}
	}

	return objs, nil
}

// 获取1000个以下对象时使用该函数
func numObjs(bucket *oss.Bucket, c *cli.Context, marker oss.Option, num int) (oss.ListObjectsResult, error) {
	return bucket.ListObjects(
		marker,
		oss.Prefix(c.String("prefix")),
		oss.MaxKeys(num),
	)
}

func printObjects(c *cli.Context, objects []oss.ObjectProperties) {

	if !c.Bool("q") {
		fmt.Println(color.New(color.FgHiCyan).Sprint("    index   size  object"))

		for index, obj := range objects {
			fmt.Printf(
				"%s  %s\n",
				color.New(color.FgHiBlack).SprintfFunc()(
					"%s",
					fmt.Sprintf("%8d%8s", index, "["+display.SmartSize(obj.Size)+"]"),
				),
				obj.Key,
			)
		}

		return
	}

	for _, obj := range objects {
		fmt.Printf("%s\n", obj.Key)
	}

	return
}
