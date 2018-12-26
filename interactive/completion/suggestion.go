package completion

import (
	"github.com/c-bata/go-prompt"
)

const (
	SysPrefix   = "!"
	flagPrefix  = "-"
	lFlagPrefix = "--"
)

var globalFlagSuggests = []prompt.Suggest{
	{Text: "--accessKey", Description: "aliyun oss accessKey"},
	{Text: "--secretKey", Description: "aliyun oss secretKey"},
	{Text: "--help", Description: "show help"},
}

var ossSubCommands = []prompt.Suggest{
	{Text: "upload", Description: "upload files to oss"},
	{Text: "download", Description: "download file from oss"},
	{Text: "delete", Description: "delete objects from oss"},
	{Text: "create", Description: "create a oss bucket"},
	{Text: "help", Description: "show help for this program"},
	{Text: "exit", Description: "exit this program"},
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
}
