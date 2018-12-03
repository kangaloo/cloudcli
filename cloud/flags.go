package cloud

import "github.com/urfave/cli"

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

// 必须二选一的两个参数一定是冲突参数

// 检查是否同时提供了会冲突的参数
func ConflictCheck(c *cli.Context, pairs [][]string) error {
	for _, pair := range pairs {
		if c.IsSet(pair[0]) && c.IsSet(pair[1]) {
			_ = cli.ShowSubcommandHelp(c)
			return NewConflictFlagErr(pair[0], pair[1])
		}
	}

	return nil
}

// 必须二选一的参数
func EitherOrCheck(c *cli.Context, pairs [][]string) error {
	for _, pair := range pairs {
		if !(c.IsSet(pair[0]) || c.IsSet(pair[1])) || (c.IsSet(pair[0]) && c.IsSet(pair[1])) {
			_ = cli.ShowSubcommandHelp(c)
			return NewEitherOrFlagErr(pair[0], pair[1])
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
