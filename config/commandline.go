package config

import (
	"errors"
	"fmt"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func parseGlobalFlags(c *cli.Context, config *Config) {

	// TODO d debug 暂时不能正常使用，需要合适的log包
	config.GlobalFlag.Debug = c.GlobalIsSet("d")

	if c.GlobalIsSet("c") {
		config.GlobalFlag.ConfigFile = c.GlobalString("c")
	}

	if c.GlobalIsSet("e") {
		config.GlobalFlag.Endpoint = c.GlobalString("e")
	}

	if c.GlobalIsSet("ak") {
		config.GlobalFlag.AccessKey = c.GlobalString("ak")
	}

	if c.GlobalIsSet("aks") {
		config.GlobalFlag.AccessKeySecret = c.GlobalString("aks")
	}
}

// 在 app的全局flags初始化后执行该函数
// 此函数必须与app结合，否则无法使用 cli.Context
func InitConfig(c *cli.Context) error {

	// create config struct
	config := &Config{}

	// parse global flags to config struct
	parseGlobalFlags(c, config)

	if !c.GlobalIsSet("c") {
		config.GlobalFlag.ConfigFile = DefaultFile()
	}

	// parse config from file to config struct
	if err := config.Parse(); err != nil {
		fmt.Println(color.New(color.FgRed).SprintfFunc()("Parsing config file failed, %s", err))
	}

	/*
		// TODO init log
		if config.debug {
			...
		}
	*/

	//
	c.App.Metadata["config"] = config
	return nil
}

func ConvertConfig(config interface{}) (*Config, error) {
	t, ok := config.(*Config)
	if ok {
		return t, nil
	}

	return nil, errors.New("convert config from interface{} to *config{} failed")
}
