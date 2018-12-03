package cloud

import "github.com/urfave/cli"

// 检查必要的参数是否都已经提供
func NecessaryCheck(c *cli.Context, flags ...string) error {
	for _, flag := range flags {
		if !c.IsSet(flag) {
			return NewNecessaryFlagErr(flag)
		}
	}

	return nil
}

// 检查是否同时提供了会冲突的参数
func ConflictCheck(c *cli.Context, pairs [][]string) error {
	for _, pair := range pairs {
		if c.IsSet(pair[0]) && c.IsSet(pair[1]) {
			return NewConflictFlagErr(pair[0], pair[1])
		}
	}

	return nil
}

// 必须二选一的参数
func EitherOrCheck(c *cli.Context, pairs [][]string) error {
	for _, pair := range pairs {
		if !(c.IsSet(pair[0]) || c.IsSet(pair[1])) || (c.IsSet(pair[0]) && c.IsSet(pair[1])) {
			return NewEitherOrFlagErr(pair[0], pair[1])
		}
	}

	return nil
}

// 检查检查参数列表中的参数是否都是定义过的，属于程序的自检测
// names 可通过 c.Names() 获取
func DefinedCheck(names []string, flagsSlice ...[]string) error {
	for _, flags := range flagsSlice {
	defined:
		for _, flag := range flags {
			for _, name := range names {
				if flag == name {
					break defined
				}
			}

			return NewNotDefinedFlagErr(flag)
		}
	}

	return nil
}
