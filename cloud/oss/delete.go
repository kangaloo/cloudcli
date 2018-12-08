package oss

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/kangaloo/cloudcli/commands/flagscheck"
	"github.com/kangaloo/cloudcli/display"
	"github.com/urfave/cli"
)

func DelBucket(c *cli.Context) error {

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

	if exist, err = client.IsBucketExist(c.String("b")); err != nil {
		return err
	}

	if !exist {
		fmt.Printf("%s bucket '%s' not exist\n", display.HiBlack("message:"), c.String("b"))
		return nil
	}

	if err = client.DeleteBucket(c.String("b")); err != nil {
		return err
	}

	fmt.Printf("%s bucket '%s' deleted\n", display.HiBlack("message:"), c.String("b"))
	return nil
}

// delete 可以完全复用 list 的 flag，删除list的结果即可
// --prefix
func DelObject(c *cli.Context) error {
	// -b -o -n --prefix --all
	// --del 真的删除，不指定该参数时，只打印将要删除的对象

	var (
		client *oss.Client
		bucket *oss.Bucket
		err    error
	)

	necessary := []string{"b"}
	atMostOne := [][]string{
		{"o", "n", "all"},
	}
	conflict := [][]string{
		{"o", "prefix"},
	}
	atLeastOne := [][]string{
		{"n", "all", "prefix", "o"},
	}

	//optional  := []string{"all"}

	if err = flagscheck.NecessaryCheck(c, necessary...); err != nil {
		return err
	}

	if err = flagscheck.AtMostOneCheck(c, atMostOne); err != nil {
		return err
	}

	if err = flagscheck.AtLeastOneCheck(c, atLeastOne); err != nil {
		return err
	}

	if err = flagscheck.ConflictCheck(c, conflict); err != nil {
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

	if err := deleteAll(bucket, c); err != nil {
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
// 该函数依赖 getAllObjs() 需要设置 --all 的默认之为true才能生效
func deleteAll(bucket *oss.Bucket, c *cli.Context) error {

	objects, err := AllObjs(bucket, c)
	if err != nil {
		return err
	}

	if !c.IsSet("del") {
		fmt.Printf("\n%s delete the listed objects with flag '--del'\n\n", display.HiBlack("message:"))
		printObjects(c, objects)
		return nil
	}

	for _, obj := range objects {
		err := deleteOne(bucket, obj.Key)
		if err != nil {
			return err
		}
	}

	return nil
}
