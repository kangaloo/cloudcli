package completion

import (
	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"os"
	"strings"
)

// todo 添加一个文件系统级别的ls命令
//  !cd 命令的处理使用 bash -c "cd" 还是 os.ChDir()
//  !pwd bash -c "pwd" / os.GetWd()

// todo 根据cli提供的内容自动生成补全提示

func Completer(d prompt.Document) []prompt.Suggest {
	if d.TextBeforeCursor() == "" {
		return []prompt.Suggest{}
	}
	args := strings.Split(d.TextBeforeCursor(), " ")

	// If PIPE is in text before the cursor, returns empty suggestions.
	for i := range args {
		if args[i] == "|" {
			return []prompt.Suggest{}
		}
	}

	if len(args) == 2 {
		if args[0] == SysPrefix+"cd" {
			return pathCompleter(d)
		}

		if args[0] == SysPrefix+"ls" {
			com := &completer.FilePathCompleter{}
			return com.Complete(d)
		}
	}

	w := d.GetWordBeforeCursor()

	// 第一级
	// todo 是否在该函数内部处理掉 args 长度检查的问题，执行底层函数出入单个字符串而非切片？
	//  个别功能复杂的函数依然需要整个切片
	if len(args) == 1 {

		// 是否为全局参数
		if strings.HasPrefix(w, "-") {
			return globalFlagCompleter(args)
		}

		// 是否为系统命令 使用!触发 e.g, !ls -> bash -c "ls"
		if strings.HasPrefix(w, SysPrefix) {
			return sysCompleter(args)
		}

		return subCommandCompleter(args)
	}

	// todo 输入空格时的提示
	// todo 显示所有可用命令
	// 输入 "." 来执行一些系统命令，如  .ls/ . ls
	// .set accessKey=LTAIpsCFIAfK0urZ

	// 长命令行参数
	if strings.HasPrefix(w, lFlagPrefix) {
		return longFlagCompleter()
	}

	// If word before the cursor starts with "-", returns CLI flag options.
	if strings.HasPrefix(w, flagPrefix) {
		return flagCompleter()
	}

	if len(args) >= 2 {

		// 文件系统匹配
		if args[len(args)-2] == "-f" {
			return fileSystemCompleter(d)
		}
	}

	return nil

}

func longFlagCompleter() []prompt.Suggest {
	return []prompt.Suggest{
		{Text: "--help", Description: "show help"},
		{Text: "--version", Description: "show version"},
	}
}

func flagCompleter() []prompt.Suggest {
	return []prompt.Suggest{
		{Text: "-h", Description: "show help"},
		{Text: "-v", Description: "show version"},
	}
}

// sysCompleter return system command and internal command
func sysCompleter(args []string) []prompt.Suggest {
	return prompt.FilterHasPrefix(
		append(SysCommands, InternalCommands...),
		args[0],
		true,
	)
}

func globalFlagCompleter(args []string) []prompt.Suggest {
	return prompt.FilterHasPrefix(globalFlagSuggests, args[len(args)-1], true)
}

func subCommandCompleter(args []string) []prompt.Suggest {
	return prompt.FilterHasPrefix(ossSubCommands, args[0], true)
}

func pathCompleter(d prompt.Document) []prompt.Suggest {
	pCompleter := &completer.FilePathCompleter{}
	pCompleter.Filter = func(fi os.FileInfo) bool {
		return fi.IsDir()
	}
	return pCompleter.Complete(d)
}

func fileSystemCompleter(d prompt.Document) []prompt.Suggest {
	fCompleter := &completer.FilePathCompleter{}
	return fCompleter.Complete(d)
}
