package executor

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/kangaloo/cloudcli/interactive/completion"
	"os"
	"os/exec"
	"strings"
)

// prompt prefix number
var Times = 0

// todo 合法命令过滤器
// todo 检查是否是支持的系统命令，防止执行不支持的系统命令
// Executor
func Executor(s string) {
	// prompt prefix number
	Times++

	s = strings.TrimSpace(s)
	if s == "" {
		return
	} else if s == "quit" || s == "exit" {
		fmt.Println("Bye!")
		os.Exit(0)
		return
	}

	// todo 检查参数是否为interactive

	// execute system command or internal command
	if strings.HasPrefix(s, completion.SysPrefix) {

		// 执行内置命令
		if supportedCMDCheck(s, completion.InternalCommands) {
			internalExecutor(s)
			return
		}

		// 执行系统命令
		if supportedCMDCheck(s, completion.SysCommands) {
			sysExecutor(s)
			return
		}

		// 不属于以上任何一种的命令，不支持的命令
		fmt.Printf("command not supported: %s\n", s)
		return
	}

	// execute application command
	appExecutor(s)

	/*
		args := strings.Split("oss "+s, " ")

		if err := command.App.Run(args); err != nil {
			fmt.Printf("Got error: %s\n", err.Error())
		}
	*/

	// todo exec(bash -c os.Args[0])
	//  交互命令和底层可执行文件为同一个文件，不加参数时进入交互模式，在交互模式中带着参数调用其本身，返回结果
	//  这样做的好处，可以在管道后面直接执行一些shell命令
	//  扩展，直接执行shell命令的可行性 设置shell命令白名单，白名单内的执行直接解析成shell命令执行

	return
}

// sysExecutor execute the system commands
// if s is a internal command invoke internalExecutor()
// else invoke shellExecutor
func sysExecutor(s string) {
	if strings.HasPrefix(s, "!cd ") {
		internalExecutor(s)
		return
	}

	shellExecutor(s)
}

func shellExecutor(s string) {
	s = strings.Split(s, completion.SysPrefix)[1]
	cmd := exec.Command("/bin/bash", "-c", s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
}

func internalExecutor(s string) {
	if strings.HasPrefix(s, "!cd") {
		s = strings.Split(s, " ")[1]
		if err := os.Chdir(s); err != nil {
			fmt.Printf("Got error: %s\n", err.Error())
		}
		return
	}

	if strings.HasPrefix(s, "!set") {
		s = strings.Split(s, " ")[1]
		// todo create the set command
		return
	}

}

func appExecutor(s string) {

	// todo os.Args[0] 是 ./cloudcli 是相对路径的，如果执行了!cd 则后续命令无法正常执行
	//  在上传相对路径的时候，还有一个 os.ChDir的动作
	//  进入交互模式后，当前目录变成了一个很重要的环境变量，需要注意，更改之后应该在改回来
	//  不改回来是否有其他影响，如读取配置文件，重新读取配置文件，
	//  配置文件的读取是相对路径还是绝对路径，相对文件还是相对执行路径的
	cmd := exec.Command("/bin/bash", "-c", os.Args[0]+" "+s)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Got error: %s\n", err.Error())
	}
}

func supportedCMDCheck(s string, commands []prompt.Suggest) bool {
	if len(strings.Split(s, " ")) > 1 {
		s = strings.Split(s, " ")[0]
	}

	for _, cmd := range commands {
		if s == cmd.Text {
			return true
		}
	}

	return false
}
