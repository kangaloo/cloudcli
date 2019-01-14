package interactive

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"github.com/kangaloo/cloudcli/interactive/completion"
	"github.com/kangaloo/cloudcli/interactive/environment"
	"github.com/kangaloo/cloudcli/interactive/executor"
	"github.com/urfave/cli"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

func Run(c *cli.Context) error {

	log.SetOutput(os.Stderr)

	var (
		binary string
		err    error
	)

	if binary, err = getBinary(); err != nil {
		return err
	}

	executor.Binary = binary
	executor.Envs = environment.NewEnvs("endpoint", "accessKey", "accessKeySecret")
	completion.EnvSuggests = []prompt.Suggest{
		{Text: "endpoint=", Description: "AliYun OSS endpoint"},
		{Text: "accessKey=", Description: "AliYun OSS accessKey"},
		{Text: "accessKeySecret=", Description: "AliYun OSS accessKeySecret"},
	}

	fmt.Printf("osscli %s\n", c.App.Version)
	fmt.Println("Please use `!` to execute system command \nand use `exit` or `Ctrl-D` to exit this program.")
	defer fmt.Println("Bye!")

	p := prompt.New(
		executor.Executor,
		completion.Completer,
		prompt.OptionTitle("OssCli"),
		prompt.OptionInputTextColor(prompt.Yellow),
		prompt.OptionCompletionWordSeparator(completer.FilePathCompletionSeparator),
		prompt.OptionPrefixTextColor(prompt.DarkGray),
		prompt.OptionSuggestionTextColor(prompt.Black),
		prompt.OptionSuggestionBGColor(prompt.White),
		prompt.OptionDescriptionBGColor(prompt.DarkGray),
		prompt.OptionSelectedSuggestionTextColor(prompt.White),
		prompt.OptionSelectedSuggestionBGColor(prompt.DarkBlue),
		prompt.OptionSelectedDescriptionBGColor(prompt.DarkBlue),
		prompt.OptionMaxSuggestion(20),
		prompt.OptionLivePrefix(livePrefix),
	)

	p.Run()

	return nil
}

// todo 在prefix里拼接换行符后，每一次键盘输入，都会换一次行
func livePrefix() (prefix string, useLivePrefix bool) {
	prefix = strconv.Itoa(executor.Times)
	pwd, err := os.Getwd()
	useLivePrefix = true
	if err != nil {
		prefix = prefix + " $ "
		return
	}

	prefix = prefix + " " + pwd + " $ "

	u, err := user.Current()
	if err != nil {
		return
	}

	prefix = strings.Replace(prefix, u.HomeDir, "~", -1)
	return
}

func getBinary() (string, error) {

	l := strings.Split(os.Args[0], string(os.PathSeparator))
	bin := l[len(l)-1]

	root, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		return "", err
	}

	bin = fmt.Sprintf("%s%s%s", root, string(os.PathSeparator), bin)

	return bin, nil
}
