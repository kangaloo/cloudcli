package interactive

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"github.com/c-bata/go-prompt/completer"
	"github.com/kangaloo/cloudcli/interactive/completion"
	"github.com/kangaloo/cloudcli/interactive/executor"
	"github.com/urfave/cli"
	"os"
	"os/user"
	"strconv"
	"strings"
)

func Run(c *cli.Context) error {

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
