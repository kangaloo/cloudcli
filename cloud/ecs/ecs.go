package ecs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/kangaloo/cloudcli/commands/flagscheck"
	"github.com/kangaloo/cloudcli/config"
	"github.com/urfave/cli"
	"os"
	"strconv"
	"text/template"
)

/*
func Ecs(c *cli.Context) error {
	fmt.Println("invoke ecs function")
	if !c.IsSet("i") {
		return cli.ShowSubcommandHelp(c)
	}
	return nil
}


func ListEcs(c *cli.Context) error {
	fmt.Println("invoke ecs function")
	if !c.IsSet("i") {
		return cli.ShowSubcommandHelp(c)
	}
	return nil
}



func Info(c *cli.Context) error {
	fmt.Println("invoke ecs > info function")

	if !c.IsSet("id") {
		return cli.ShowCommandHelp(c, c.Command.Name)
	}

	fmt.Println(c.String("id"))
	return nil
}

*/

func extractParameter(confer interface{}) (ak, aks, region, endpoint, temp string, err error) {
	var conf *config.Config

	if conf, err = config.ConvertConfig(confer); err != nil {
		return
	}

	// set accessKey
	ak = conf.Global.AccessKey
	if len(conf.Ecs.AccessKey) != 0 {
		ak = conf.Ecs.AccessKey
	}
	if len(conf.GlobalFlag.AccessKey) != 0 {
		ak = conf.GlobalFlag.AccessKey
	}

	// set accessKeySecret
	aks = conf.Global.AccessKeySecret
	if len(conf.Ecs.AccessKeySecret) != 0 {
		aks = conf.Ecs.AccessKeySecret
	}
	if len(conf.GlobalFlag.AccessKeySecret) != 0 {
		aks = conf.GlobalFlag.AccessKeySecret
	}

	// set regionID
	region = conf.Ecs.RegionID

	// set endpoint
	endpoint = conf.Ecs.Endpoint
	if len(conf.GlobalFlag.Endpoint) != 0 {
		endpoint = conf.GlobalFlag.Endpoint
	}

	// todo 设置由命令行参数指定的template
	temp = conf.Ecs.Format

	if err = flagscheck.LengthCheck(ak, aks, region); err != nil {
		return
	}

	return
}

func ListEcs(c *cli.Context) error {

	ak, aks, region, endpoint, temp, err := extractParameter(c.App.Metadata["config"])
	if err != nil {
		return err
	}

	client, err := ecs.NewClientWithAccessKey(region, ak, aks)
	if err != nil {
		return err
	}

	request := ecs.CreateDescribeInstancesRequest()
	request.SetDomain(endpoint)

	instances, err := getAllEcs(client, request)
	if err != nil {
		return err
	}

	if err := handleResp(instances, temp); err != nil {
		return err
	}

	return nil
}

func getAllEcs(c *ecs.Client, r *ecs.DescribeInstancesRequest) ([]ecs.Instance, error) {

	instances := make([]ecs.Instance, 0, 1024)
	page := requests.Integer("1")
	ps, _ := r.PageSize.GetValue()

	for {
		r.PageNumber = page
		resp, err := c.DescribeInstances(r)
		if err != nil {
			return nil, err
		}

		if len(resp.Instances.Instance) == 0 {
			break
		}

		instances = append(instances, resp.Instances.Instance...)

		if len(resp.Instances.Instance) < ps {
			break
		}

		n, err := page.GetValue()
		if err != nil {
			return nil, err
		}

		n += 1
		page = requests.Integer(strconv.Itoa(n))
	}

	return instances, nil
}

func handleResp(instances []ecs.Instance, temp string) error {

	if len(temp) == 0 {
		for _, instance := range instances {
			if err := formatToJson(&instance); err != nil {
				return err
			}
		}

		return nil
	}

	// format output with template string
	fmt.Println(temp)

	t := template.New("ecs output")
	if _, err := t.Parse(temp); err != nil {
		return err
	}

	for _, instance := range instances {
		if err := t.Execute(os.Stdout, instance); err != nil {
			return err
		}
		fmt.Println()
	}

	return nil
}

func formatToJson(ecs *ecs.Instance) error {
	i, err := json.Marshal(ecs)
	if err != nil {
		return err
	}

	dst := bytes.NewBuffer(make([]byte, 0, 1024))
	if err := json.Indent(dst, i, "", "\t"); err != nil {
		return err
	}

	fmt.Printf("------------------ %s ----------------\n", ecs.InstanceName)
	fmt.Println(dst)
	fmt.Println()
	return nil
}
