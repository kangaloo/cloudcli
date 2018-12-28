package completion

import (
	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"os"
	"strings"
)

//  todo !cd 命令的处理使用 bash -c "cd" 还是 os.ChDir()
//   !pwd bash -c "pwd" / os.GetWd()

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

	w := d.GetWordBeforeCursor()

	// 第一级命令 e.g. oss|help | 系统命令 e.g. cd|ls|pwd
	// !set 命令用于设置全局参数
	if len(args) == 1 {

		// 是否为系统命令 使用!触发 e.g, !ls -l -> `bash -c "ls -l"`
		if strings.HasPrefix(w, SysPrefix) {
			return sysCompleter(args, d)
		}

		return appCompleter(args, d)
	}

	// 第二级命令 e.g. upload download
	if len(args) == 2 {

		if strings.HasPrefix(args[0], SysPrefix) {
			return sysCompleter(args, d)
		}

		return appCompleter(args, d)
	}

	if len(args) > 2 {
		return appCompleter(args, d)
	}

	// 第一级
	// todo 是否在该函数内部处理掉 args 长度检查的问题，执行底层函数出入单个字符串而非切片？
	//  个别功能复杂的函数依然需要整个切片

	// todo 输入空格时的提示
	// todo 显示所有可用命令

	return nil
}

func appCompleter(args []string, d prompt.Document) []prompt.Suggest {

	if len(args) == 1 {
		return subCommandCompleter(args)
	}

	if len(args) >= 2 {
		switch args[0] {
		case "oss":
			return ossSubCommandCompleter(args, d)
		case "help":
			return nil
		case "interactive":
			return nil
		default:
			return nil
		}
	}

	return nil
}

func ossSubCommandCompleter(args []string, d prompt.Document) []prompt.Suggest {
	if len(args) == 2 {
		return prompt.FilterHasPrefix(OssSubCommands, args[1], true)
	}

	switch args[1] {
	case "upload", "ul":
		return ossUploadCompleter(args, d)
	case "download", "dl":
		return ossDownloadCompleter(args, d)
	case "list", "ls":
		return ossListCompleter(args)
	case "list_bucket", "lsbk":
		return nil
	case "create", "ct":
		return ossCreateCompleter(args)
	case "delete", "del":
		return ossDelCompleter(args)
	case "delete_bucket", "delbk":
		return ossCreateCompleter(args)
	}

	return nil
}

// sysCompleter return system commands and internal commands
func sysCompleter(args []string, d prompt.Document) []prompt.Suggest {

	if len(args) == 2 {
		if args[0] == SysPrefix+"cd" {
			return pathCompleter(d)
		}

		if args[0] == SysPrefix+"ls" {
			return fileSystemCompleter(d)
		}

		if args[0] == SysPrefix+"set" {
			return setCompleter(args)
		}
	}

	if len(args) == 1 {
		return prompt.FilterHasPrefix(
			append(SysCommands, InternalCommands...),
			args[0],
			true,
		)
	}

	return nil
}

func subCommandCompleter(args []string) []prompt.Suggest {
	return prompt.FilterHasPrefix(SubCommands, args[0], true)
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

func setCompleter(args []string) []prompt.Suggest {
	return nil
}
