package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/kangaloo/cloudcli/commands/flagscheck"
	"github.com/kangaloo/cloudcli/display"
	"github.com/kangaloo/cloudcli/file"
	"github.com/urfave/cli"
	"log"
	"os"
	"strings"
)

// --prefix
func Upload(c *cli.Context) error {

	// check flags
	{
		var (
			necessary  = []string{"b"}
			conflict   = [][]string{{"R", "r"}, {"o", "prefix"}}
			atLeastOne = [][]string{{"r", "R", "f"}}
			err        error
		)

		if err = flagscheck.NecessaryCheck(c, necessary...); err != nil {
			return err
		}

		if err = flagscheck.ConflictCheck(c, conflict); err != nil {
			return err
		}

		if err = flagscheck.AtLeastOneCheck(c, atLeastOne); err != nil {
			return err
		}
	}

	var (
		client *oss.Client
		bucket *oss.Bucket
		err    error
	)

	// Create client
	if client, err = NewOssClient(c.App.Metadata["config"]); err != nil {
		return err
	}

	// Get bucket
	if bucket, err = client.Bucket(c.String("b")); err != nil {
		return err
	}

	// Upload files
	if err = upload(bucket, c); err != nil {
		return err
	}

	return nil
}

func upload(bucket *oss.Bucket, c *cli.Context) error {

	var (
		root       string
		prefix     string
		singleFile string
		overwrite  bool
		err        error
	)

	// parse flags
	root = c.String("r")
	prefix = c.String("prefix")
	overwrite = c.Bool("overwrite")

	// 上传单个文件
	// 此处可根据是否提供了-o参数判断
	if !(c.Bool("R") || c.IsSet("r")) {

		singleFile = c.String("f")
		object := addPrefix(singleFile, prefix)

		if err = uploadOneFile(bucket, singleFile, object, overwrite); err != nil {
			return err
		}

		return nil
	}

	// 上传目录，两种情况 -R / -r
	// -R path = "."
	// -r path = specified_path
	if c.Bool("R") {
		if err = uploadRecursively(bucket, "", prefix, overwrite); err != nil {
			return err
		}

		return nil
	}

	if err = uploadRecursively(bucket, root, prefix, overwrite); err != nil {
		return err
	}

	return nil
}

// 核心上传函数
func uploadOneFile(bucket *oss.Bucket, file, object string, overwrite bool) error {

	var (
		exist bool
		err   error
	)

	if exist, err = bucket.IsObjectExist(object); err != nil {
		return err
	}

	if !exist {
		log.Printf("start upload object [ %s ]", display.Green("%s", object))
		if err = bucket.PutObjectFromFile(object, file, oss.Progress(&ossProgressListener{})); err != nil {
			return err
		}

		return nil
	}

	if !overwrite {
		log.Printf("skip object [ %s ] cause already exist", display.Yellow("%s", object))
		return nil
	}

	log.Printf("overwrite object [ %s ]", display.Yellow("%s", object))
	if err = bucket.PutObjectFromFile(object, file, oss.Progress(&ossProgressListener{})); err != nil {
		return err
	}

	return nil
}

// 批量上传函数
// 需要上传当前目录的内容时，root参数传空字符串
// 不需要prefix参数时，传空字符串
func uploadRecursively(bucket *oss.Bucket, root, prefix string, overwrite bool) error {

	var (
		files []string
		cwd   string
		err   error
	)

	if root != "" {
		if err = os.Chdir(root); err != nil {
			return err
		}
	}

	if cwd, err = os.Getwd(); err != nil {
		return err
	}

	if err = file.CollectFiles(&files, cwd); err != nil {
		return err
	}

	if prefix == "" {
		for _, f := range files {
			if err = uploadOneFile(bucket, f, f, overwrite); err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range files {
		object := addPrefix(f, prefix)
		if err = uploadOneFile(bucket, f, object, overwrite); err != nil {
			return err
		}
	}

	return nil
}

func addPrefix(path, prefix string) string {
	if strings.HasSuffix(prefix, "/") {
		return fmt.Sprintf("%s%s", prefix, path)
	}
	return fmt.Sprintf("%s%s%s", prefix, string(os.PathSeparator), path)
}
