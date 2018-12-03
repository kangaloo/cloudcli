package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/fatih/color"
	"github.com/kangaloo/cloudcli/config"
)

func NewOssClient(configInter interface{}) (*oss.Client, error) {

	var (
		endpoint, ak ,aks string
	)

	conf, err := config.ConvertConfig(configInter)

	if err != nil {
		return nil, err
	}

	endpoint = conf.Oss.Endpoint

	if len(conf.GlobalFlag.Endpoint) > 0 {
		endpoint = conf.GlobalFlag.Endpoint
	}

	ak  = conf.Global.AccessKey

	if len(conf.Oss.AccessKey) >0 {
		ak = conf.Oss.AccessKey
	}

	if len(conf.GlobalFlag.AccessKey) > 0 {
		ak = conf.GlobalFlag.AccessKey
	}

	aks = conf.Global.AccessKeySecret

	if len(conf.Oss.AccessKeySecret) > 0 {
		aks = conf.Oss.AccessKeySecret
	}

	if len(conf.GlobalFlag.AccessKeySecret) > 0 {
		aks = conf.GlobalFlag.AccessKeySecret
	}

	return oss.New(endpoint, ak, aks)
}

// OssProgressListener show status
type ossProgressListener struct{}

// ProgressChanged show the status
func (listener *ossProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		fmt.Printf("Transfer Started, ConsumedBytes: %d, TotalBytes %d (%s)\n",
			event.ConsumedBytes, event.TotalBytes, smartSize(event.TotalBytes))
	case oss.TransferDataEvent:
		fmt.Printf(
			"\rTransfer Data, ConsumedBytes: %s, TotalBytes %s, %d%%.",
			color.New(color.FgGreen).SprintfFunc()("%d", event.ConsumedBytes),
			color.New(color.FgGreen).SprintfFunc()("%d", event.TotalBytes),
			event.ConsumedBytes*100/event.TotalBytes,
			)
	case oss.TransferCompletedEvent:
		fmt.Printf("\nTransfer Completed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferFailedEvent:
		fmt.Printf("\nTransfer Failed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	default:
	}
}

func smartSize(s int64) string {

	if s < 1024 {
		return fmt.Sprintf("%d B", s)
	}

	if s < 1024*1024 {
		return fmt.Sprintf("%d KB", s/1024)
	}

	if s < 1024*1024*1024 {
		return fmt.Sprintf("%d MB", s/1024/1024)
	}

	if s < 1024*1024*1024*1024 {
		return fmt.Sprintf("%d GB", s/1024/1024/1024)
	}

	return fmt.Sprintf("%d TB", s/1024/1024/1024/1024)
}