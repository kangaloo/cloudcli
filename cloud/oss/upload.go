package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/kangaloo/cloudcli/display"
	"github.com/kangaloo/cloudcli/file"
	"github.com/urfave/cli"
	"log"
)

// --prefix
func Upload(c *cli.Context) error {
	// 参数检查
	// 上传单个文件时是否支持 --prefix 参数
	// -r 参数不能为空字符串
	// -R / -r 不能同时出现
	// -b 必须参数

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
	// todo 分层，分成参数处理层和底层api调用层，底层函数直接接受各种参数
	// 该函数负责解析参数

	var (
		files []string
		path  string
		err   error
	)

	// 上传单个文件
	if !(c.Bool("R") || c.IsSet("r")) {
		// object = prefix + file
		if err = uploadOneFile(bucket, "", "", c.Bool("overwrite")); err != nil {
			return err
		}

		return nil
	}

	// 上传目录，两种情况 -R / -r
	// -R path = "."
	// -r path = specified_path

	if err = file.CollectFiles(&files, path, ""); err != nil {
		return err
	}

	if err = uploadRecursively(bucket, files, "", false); err != nil {
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

func uploadRecursively(bucket *oss.Bucket, files []string, prefix string, overwrite bool) error {

	return nil
}
