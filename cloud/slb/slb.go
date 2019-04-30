package slb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/slb"
	"github.com/kangaloo/cloudcli/commands/flagscheck"
	"github.com/kangaloo/cloudcli/config"
	"github.com/urfave/cli"
	"os"
	"text/template"
)

func extractParameter(confer interface{}) (ak, aks, region, endpoint, temp string, err error) {
	var conf *config.Config

	if conf, err = config.ConvertConfig(confer); err != nil {
		return
	}

	// set accessKey
	ak = conf.Global.AccessKey
	if len(conf.Slb.AccessKey) != 0 {
		ak = conf.Slb.AccessKey
	}
	if len(conf.GlobalFlag.AccessKey) != 0 {
		ak = conf.GlobalFlag.AccessKey
	}

	// set accessKeySecret
	aks = conf.Global.AccessKeySecret
	if len(conf.Slb.AccessKeySecret) != 0 {
		aks = conf.Slb.AccessKeySecret
	}
	if len(conf.GlobalFlag.AccessKeySecret) != 0 {
		aks = conf.GlobalFlag.AccessKeySecret
	}

	// set regionID
	region = conf.Slb.RegionID

	// set endpoint
	endpoint = conf.Slb.Endpoint
	if len(conf.GlobalFlag.Endpoint) != 0 {
		endpoint = conf.GlobalFlag.Endpoint
	}

	// todo 设置由命令行参数指定的template
	temp = conf.Slb.Format

	if err = flagscheck.LengthCheck(ak, aks, region); err != nil {
		return
	}

	return
}

// 主要负责处理命令行参数
func ListLB(c *cli.Context) error {

	var (
		ak, aks, region, endpoint, temp string
		client                          *slb.Client
		request                         *slb.DescribeLoadBalancersRequest
		response                        *slb.DescribeLoadBalancersResponse
		err                             error
	)

	if ak, aks, region, endpoint, temp, err = extractParameter(c.App.Metadata["config"]); err != nil {
		return err
	}

	if client, err = slb.NewClientWithAccessKey(region, ak, aks); err != nil {
		return err
	}

	request = slb.CreateDescribeLoadBalancersRequest()
	request.SetDomain(endpoint)

	// 发起查询请求
	if response, err = client.DescribeLoadBalancers(request); err != nil {
		return err
	}

	// 处理结果和参数
	if err = handleResp(response.LoadBalancers.LoadBalancer, temp); err != nil {
		return err
	}

	return nil
}

// 负责发起请求 格式化输出等
func DescribeLB(c *cli.Context) error {

	necessary := []string{"i"}
	if err := flagscheck.NecessaryCheck(c, necessary...); err != nil {
		return err
	}

	ak, aks, region, endpoint, _, err := extractParameter(c.App.Metadata["config"])
	if err != nil {
		return err
	}

	client, err := slb.NewClientWithAccessKey(region, ak, aks)
	if err != nil {
		return err
	}

	response, err := descLBAttr(c.String("i"), endpoint, client)
	if err != nil {
		return err
	}

	if err := FormatToJson(response); err != nil {
		return err
	}

	hresponse, err := descHealthStatus(c.String("i"), endpoint, client)
	if err != nil {
		return err
	}

	return FormatToJson(hresponse)
}

func descLBAttr(id, endpoint string, client *slb.Client) (*slb.DescribeLoadBalancerAttributeResponse, error) {
	request := slb.CreateDescribeLoadBalancerAttributeRequest()
	request.SetDomain(endpoint)
	request.LoadBalancerId = id

	return client.DescribeLoadBalancerAttribute(request)
}

func descHealthStatus(id, endpoint string, client *slb.Client) (*slb.DescribeHealthStatusResponse, error) {
	request := slb.CreateDescribeHealthStatusRequest()
	request.SetDomain(endpoint)
	request.LoadBalancerId = id

	return client.DescribeHealthStatus(request)
}

func handleResp(lbs []slb.LoadBalancer, temp string) error {

	if len(temp) == 0 {
		for _, lb := range lbs {
			if err := formatToJson(&lb); err != nil {
				return err
			}
		}

		return nil
	}

	// format output with template string
	fmt.Println(temp)

	t := template.New("slb output")
	if _, err := t.Parse(temp); err != nil {
		return err
	}

	for _, lb := range lbs {
		if err := t.Execute(os.Stdout, lb); err != nil {
			return err
		}
		fmt.Println()
	}

	return nil
}

func formatToJson(lb *slb.LoadBalancer) error {
	i, err := json.Marshal(lb)
	if err != nil {
		return err
	}

	dst := bytes.NewBuffer(make([]byte, 0, 1024))
	if err := json.Indent(dst, i, "", "\t"); err != nil {
		return err
	}

	fmt.Printf("------------------ %s ----------------\n", lb.LoadBalancerName)
	fmt.Println(dst)
	fmt.Println()
	return nil
}

func FormatToJson(lb interface{}) error {
	i, err := json.Marshal(lb)
	if err != nil {
		return err
	}

	dst := bytes.NewBuffer(make([]byte, 0, 1024))
	if err := json.Indent(dst, i, "", "\t"); err != nil {
		return err
	}

	fmt.Println(dst)
	fmt.Println()
	return nil
}
