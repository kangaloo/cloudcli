package executor

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/kangaloo/cloudcli/interactive/completion"
	"github.com/kangaloo/cloudcli/interactive/environment"
	"os"
	"os/exec"
	"strings"
)

var (
	// prompt prefix number
	Times  = 0
	Binary string
	Envs   []environment.Env
)

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
			shellExecutor(s)
			return
		}

		// 不属于以上任何一种的命令，不支持的命令
		fmt.Printf("command not supported: %s\n", s)
		return
	}

	// execute application command
	appExecutor(s)

	// todo exec(bash -c os.Args[0])
	//  交互命令和底层可执行文件为同一个文件，不加参数时进入交互模式，在交互模式中带着参数调用其本身，返回结果
	//  这样做的好处，可以在管道后面直接执行一些shell命令
	//  扩展，直接执行shell命令的可行性 设置shell命令白名单，白名单内的执行直接解析成shell命令执行

	return
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

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if strings.HasPrefix(s, "!cd") {
		s = strings.Split(s, " ")[1]
		if err := os.Chdir(s); err != nil {
			fmt.Printf("Got error: %s\n", err.Error())
		}
		return
	}

	if strings.HasPrefix(s, "!set") {
		s = strings.Split(s, " ")[1]
		setEnv(s)
		return
	}

	if strings.HasPrefix(s, "!env") {
		for _, env := range Envs {
			env.Display()
		}
		return
	}
}

func appExecutor(s string) {

	// todo
	//  在上传相对路径的时候，还有一个 os.ChDir的动作
	//  进入交互模式后，当前目录变成了一个很重要的环境变量，需要注意，更改之后应该在改回来
	//  不改回来是否有其他影响，如读取配置文件，重新读取配置文件，
	//  配置文件的读取是相对路径还是绝对路径，相对文件还是相对执行路径的

	// todo 获取set命令设置的env，加入到全局参数里

	// todo 增加全局参数 ak aks endpoint

	globalFlags := generateFlags(Envs)

	cmd := exec.Command("/bin/bash", "-c", Binary+" "+globalFlags+" "+s)
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

// 生成全局命令行参数
func generateFlags(envs []environment.Env) string {

	var flags string

	for _, env := range envs {

		if env.Value == "" {
			continue
		}

		flags = flags + env.Flag + " " + env.Value + " "
	}

	return flags
}

func setEnv(s string) {
	t := strings.Split(s, "=")
	for index, env := range Envs {
		if env.Key == t[0] {
			Envs[index].Value = t[1]
		}
	}
}
