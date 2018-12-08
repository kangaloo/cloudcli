package flagscheck

import (
	"errors"
	"github.com/urfave/cli"
)

// 至多有一个的参数
// 此函数可取代 ConflictCheck
func AtMostOneCheck(c *cli.Context, lists [][]string) error {
	for _, list := range lists {
		var provided []string

		for _, flag := range list {
			if c.IsSet(flag) {
				provided = append(provided, flag)
			}
		}

		if len(provided) > 1 {
			return NewConflictFlagErr(provided)
		}
	}
	return nil
}

func AtLeastOneCheck(c *cli.Context, lists [][]string) error {
	for _, list := range lists {
		var provided []string

		for _, flag := range list {
			if c.IsSet(flag) {
				provided = append(provided, flag)
			}
		}

		if len(provided) == 0 {
			return NewAtLeastOneErr(list)
		}
	}
	return nil
}

// 必须二选一的参数
func EitherOrCheck(c *cli.Context, pairs [][]string) error {
	for _, pair := range pairs {

		/*
			if !(c.IsSet(pair[0]) || c.IsSet(pair[1])) || (c.IsSet(pair[0]) && c.IsSet(pair[1])) {
				_ = cli.ShowSubcommandHelp(c)
				return NewEitherOrFlagErr(pair[0], pair[1])
			}
		*/

		if !(c.IsSet(pair[0]) || c.IsSet(pair[1])) {
			_ = cli.ShowSubcommandHelp(c)
			return NewEitherOrFlagErr(pair[0], pair[1])
		}

		if c.IsSet(pair[0]) && c.IsSet(pair[1]) {
			_ = cli.ShowSubcommandHelp(c)
			return NewConflictFlagErr(pair)
		}
	}

	return nil
}

// MustOnlyOne 只能且必须提供一个的参数
func MustOnlyOneCheck(c *cli.Context, lists [][]string) error {
	for _, list := range lists {
		var provided []string

		for _, flag := range list {
			if c.IsSet(flag) {
				provided = append(provided, flag)
			}
		}

		if len(provided) == 0 {
			return NewAtLeastOneErr(list)
		}

		if len(provided) > 1 {
			return NewConflictFlagErr(provided)
		}
	}
	return nil
}

// 检查是否同时提供了会冲突的参数
func ConflictCheck(c *cli.Context, pairs [][]string) error {
	for _, pair := range pairs {
		if c.IsSet(pair[0]) && c.IsSet(pair[1]) {
			_ = cli.ShowSubcommandHelp(c)
			return NewConflictFlagErr(pair)
		}
	}

	return nil
}

// 检查检查参数列表中的参数是否都是定义过的，属于程序的自检测
// 不需要，cli层自动处理
// flag provided but not defined: -p
func DefinedCheck(c *cli.Context, flagsSlice ...[]string) error {
	names := c.FlagNames()

	for _, flags := range flagsSlice {
	defined:
		for _, flag := range flags {
			for _, name := range names {
				if flag == name {
					break defined
				}
			}

			_ = cli.ShowSubcommandHelp(c)
			return NewNotDefinedFlagErr(flag)
		}
	}

	return nil
}

// 检查必要的参数是否都已经提供
func NecessaryCheck(c *cli.Context, flags ...string) error {
	for _, flag := range flags {
		if !c.IsSet(flag) {
			_ = cli.ShowSubcommandHelp(c)
			return NewNecessaryFlagErr(flag)
		}
	}

	return nil
}

func LengthCheck(params ...string) error {
	for _, param := range params {
		if param == "" {
			return errors.New("parameter is empty string")
		}
	}

	return nil
}

// Command line parameters

func checkGlobalParam() error {
	return nil
}
