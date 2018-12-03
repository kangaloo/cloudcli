package oss

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/kangaloo/cloudcli/cloud"
	"github.com/kangaloo/cloudcli/display"
	"github.com/urfave/cli"
)

func DelBucket(c *cli.Context) error {

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

	if err = client.DeleteBucket(c.String("b")); err != nil {
		return err
	}

	return nil
}

func DelObject(c *cli.Context) error {

	var (
		client *oss.Client
		bucket *oss.Bucket
		err    error
	)

	necessary := []string{"b"}
	eitherOr := [][]string{{"o", "all"}}
	//optional  := []string{"all"}

	if err = cloud.NecessaryCheck(c, necessary...); err != nil {
		return err
	}

	if err = cloud.EitherOrCheck(c, eitherOr); err != nil {
		return err
	}

	if client, err = NewOssClient(c.App.Metadata["config"]); err != nil {
		return err
	}

	if bucket, err = client.Bucket(c.String("b")); err != nil {
		return err
	}

	if c.IsSet("o") {
		if err = deleteOne(bucket, c.String("o")); err != nil {
			return err
		}
		return nil
	}

	if err := deleteAll(bucket); err != nil {
		return err
	}

	return nil
}

// deleteOne delete one object by object name
func deleteOne(bucket *oss.Bucket, key string) error {

	var (
		exist bool
		err   error
	)

	if exist, err = bucket.IsObjectExist(key); err != nil {
		return err
	}

	if !exist {
		fmt.Printf("%s object '%s' not exist\n", display.HiBlack("message:"), key)
		return nil
	}

	if err = bucket.DeleteObject(key); err != nil {
		return err
	}

	if exist, err = bucket.IsObjectExist(key); err != nil {
		return err
	}

	if exist {
		fmt.Printf("%s object '%s' is still here, delete failed\n", display.Red("message:"), key)
		return errors.New("delete failed")
	}

	fmt.Printf("%s object '%s' deleted\n", display.HiBlack("message:"), key)

	return nil
}

// deleteAll delete all objects in a bucket
func deleteAll(bucket *oss.Bucket) error {

	objects, err := listAllObjs(bucket)
	if err != nil {
		return err
	}

	for _, obj := range objects {
		err := deleteOne(bucket, obj.Key)
		if err != nil {
			return err
		}
	}

	return nil
}
