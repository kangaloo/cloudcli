package main

import (
	"fmt"
	"github.com/kangaloo/cloudcli/commands"
	"github.com/kangaloo/cloudcli/commands/ecs"
	"github.com/kangaloo/cloudcli/commands/interactivemode"
	"github.com/kangaloo/cloudcli/commands/oss"
	"github.com/kangaloo/cloudcli/commands/slb"
	"github.com/kangaloo/cloudcli/config"
	"github.com/kangaloo/cloudcli/display"
	"github.com/urfave/cli"
	"os"
)

func main() {

	// TODO 先写config包，后写log包
	//	 config包出现错误时，错误日志输出到标准输出，并退出

	app := cli.NewApp()
	app.Name = "cloudcli"
	app.Usage = "Aliyun API command line tool"
	app.Email = "lixy4@belink.com"
	app.Author = "Li Xiangyang"
	app.Version = "1.0.0"

	// 全局命令行参数
	commands.AddGlobalFlags(app, commands.GlobalFlags)

	// 初始化配置
	app.Before = config.InitConfig

	// 初始化子命令 目前支持 oss ecs slb
	commands.AddCmd(app, oss.Oss)
	commands.AddCmd(app, ecs.Ecs)
	commands.AddCmd(app, slb.Slb)
	commands.AddCmd(app, interactivemode.Interactive)

	err := app.Run(os.Args)

	if err != nil {
		fmt.Println(display.Red("ERROR: %s", err))
		os.Exit(1)
	}
}
