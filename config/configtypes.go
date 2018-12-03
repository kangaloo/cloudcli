package config

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Global `yaml:"global"`
	Oss    `yaml:"oss"`
	Ecs    `yaml:"ecs"`
	Slb    `yaml:"slb"`
	GlobalFlag
}

func (c *Config) Parse() error {
	if len(c.GlobalFlag.ConfigFile) <= 0 {
		return NotSetErr
	}

	conf, err := ioutil.ReadFile(c.GlobalFlag.ConfigFile)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(conf, c)
}

type Global struct {
	Debug           bool   `yaml:"debug"`
	AccessKey       string `yaml:"accesskey"`
	AccessKeySecret string `yaml:"accesskeysecret"`
}

type GlobalFlag struct {
	Debug           bool
	ConfigFile      string
	Endpoint        string
	AccessKey       string
	AccessKeySecret string
}

type Oss struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKey       string `yaml:"accesskey"`
	AccessKeySecret string `yaml:"accesskeysecret"`
	RegionID        string `yaml:"regionid"`
}

type Ecs struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKey       string `yaml:"accesskey"`
	AccessKeySecret string `yaml:"accesskeysecret"`
	RegionID        string `yaml:"regionid"`
	Format          string `yaml:"format"`
	ID              string
	PageSize        requests.Integer
}

type Slb struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKey       string `yaml:"accesskey"`
	AccessKeySecret string `yaml:"accesskeysecret"`
	RegionID        string `yaml:"regionid"`
	Format          string `yaml:"format"`
	ID              string
}
