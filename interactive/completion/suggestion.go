package completion

import (
	"github.com/c-bata/go-prompt"
)

const (
	SysPrefix   = "!"
	flagPrefix  = "-"
	lFlagPrefix = "--"
)

// todo 交互模式暂时不支持全局参数
//  help 子命令有completer函数生成，每个completer函数都可生成该子命令
//  set 内置命令取代全局参数 需要修改 executor 在执行时将set命令产生的env转换成全局参数

var SubCommands = []prompt.Suggest{
	{Text: "oss", Description: "AliYun OSS API tool"},
	{Text: "help", Description: "show help for this program"},
	{Text: "exit", Description: "exit this program"},
}

// todo 字命令的别名
var OssSubCommands = []prompt.Suggest{
	{Text: "upload", Description: "upload files to oss"},
	{Text: "download", Description: "download file from oss"},
	{Text: "list", Description: "list objects in oss"},
	{Text: "list_bucket", Description: "list all oss buckets"},
	{Text: "delete", Description: "delete objects from oss"},
	{Text: "delete_bucket", Description: "delete oss bucket"},
	{Text: "create", Description: "create a oss bucket"},
	{Text: "help", Description: "show help for this program"},
}

// todo 增加 cat less more tail 等查看文件内容的命令
var SysCommands = []prompt.Suggest{
	{Text: SysPrefix + "ls", Description: "list directory contents for local file system"},
	{Text: SysPrefix + "pwd", Description: "print name of current/working directory"},
	{Text: SysPrefix + "date", Description: "display or set date and time"},
	{Text: SysPrefix + "uptime", Description: "tell how long the system has been running"},
}

var InternalCommands = []prompt.Suggest{
	{Text: SysPrefix + "cd", Description: "change directory for local file system"},
	{Text: SysPrefix + "set", Description: "set env"},
	{Text: SysPrefix + "env", Description: "display env"},
}

var EnvSuggests []prompt.Suggest

var SetSubCommands = []prompt.Suggest{
	{Text: "accessKey", Description: ""},
	{Text: "secretKey", Description: ""},
}

var CommonFlags = []prompt.Suggest{
	{Text: "-h", Description: ""},
	{Text: "--help", Description: ""},
}

var CommonCommands = []prompt.Suggest{
	{Text: "help", Description: ""},
}
