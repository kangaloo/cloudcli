package cloud

import "github.com/kangaloo/cloudcli/config"

var Config = &config.Config{}

func init() {
	// 完成了默认配置文件的选择
	parseEnv(Config)

	// todo 判读是否使用默认文件
	Config.GlobalFlag.ConfigFile = config.DefaultFile()

	if err := Config.Parse(); err != nil {
		return
	}
}

// todo 解析环境变量 作为全局变量 包含默认配置文件
func parseEnv(c *config.Config) {

}
