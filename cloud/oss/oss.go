package oss

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/kangaloo/cloudcli/commands/flagscheck"
	"github.com/kangaloo/cloudcli/config"
	"github.com/kangaloo/cloudcli/display"
)

func NewOssClient(configInter interface{}) (*oss.Client, error) {

	var (
		endpoint, ak, aks string
		conf              *config.Config
		err               error
	)

	if conf, err = config.ConvertConfig(configInter); err != nil {
		return nil, err
	}

	endpoint = conf.Oss.Endpoint

	if len(conf.GlobalFlag.Endpoint) > 0 {
		endpoint = conf.GlobalFlag.Endpoint
	}

	ak = conf.Global.AccessKey

	if len(conf.Oss.AccessKey) > 0 {
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

	// if the parameters are empty strings, oss.New will panic
	if err = flagscheck.LengthCheck(endpoint, ak, aks); err != nil {
		return nil, err
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
			event.ConsumedBytes, event.TotalBytes, display.HiBlack(display.SmartSize(event.TotalBytes)))
	case oss.TransferDataEvent:
		fmt.Printf(
			"\rTransfer Data, ConsumedBytes: %s, TotalBytes %s, %s.",
			display.HiBlack("%d", event.ConsumedBytes),
			display.HiBlack("%d", event.TotalBytes),
			display.HiBlack("%d%%", event.ConsumedBytes*100/event.TotalBytes),
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
