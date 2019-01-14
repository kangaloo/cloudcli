package cloud

import (
	oss2 "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/c-bata/go-prompt"
	"github.com/kangaloo/cloudcli/cloud/oss"
	"sync"
)

var (
	bucketList = &sync.Map{}
)

func fetchBuckets() {

	client, err := oss.NewOssClient(Config)
	if err != nil {
		return
	}

	list, err := client.ListBuckets()
	if err != nil {
		return
	}

	bucketList.Store("buckets", list.Buckets)
}

func BucketCompleter() []prompt.Suggest {

	go fetchBuckets()

	l, ok := bucketList.Load("buckets")
	if !ok {
		return []prompt.Suggest{}
	}

	list, ok := l.([]oss2.BucketProperties)

	if !ok {
		return []prompt.Suggest{}
	}

	if len(list) == 0 {
		return []prompt.Suggest{}
	}

	suggests := make([]prompt.Suggest, len(list))

	for index, suggest := range list {
		suggests[index] = prompt.Suggest{Text: suggest.Name}
	}

	return suggests
}
