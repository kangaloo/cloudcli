package oss

import (
	"github.com/kangaloo/cloudcli/cloud"
	"github.com/urfave/cli"
)

func DelBucket(c *cli.Context) error  {
	// TODO 参数检查

	necessary := []string{"b"}

	if err := cloud.NecessaryCheck(c, necessary...); err != nil {
		return err
	}

	client, err := NewOssClient(c.App.Metadata["config"])

	if err != nil {
		return err
	}

	err = client.DeleteBucket(c.String("b"))

	if err != nil {
		return err
	}

	return nil
}

func DelObject(c *cli.Context) error {

	// TODO 参数检查

	necessary := []string{"b", "o"}

	if err := cloud.NecessaryCheck(c, necessary...); err != nil {
		return err
	}

	client, err := NewOssClient(c.App.Metadata["config"])

	if err != nil {
		return err
	}

	bucket, err := client.Bucket(c.String("b"))

	if err != nil {
		return err
	}

	err = bucket.DeleteObject(c.String("o"))
	if err != nil {
		return err
	}

	return nil
}

// deleteOne delete one object by object name
func deleteOne()  {

}

// deleteAll delete all objects in a bucket
func deleteAll()  {

}
