package file

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func IsFileExist(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		return os.IsExist(err)
	}

	return true
}

func IsDir(file string) bool {
	fi, err := os.Stat(file)
	if err != nil {
		return false
	}

	return fi.IsDir()
}

func DefaultConfDir() string {
	// 相对于当前工作目录的可执行文件所在目录
	base := filepath.Dir(os.Args[0])

	// configuration file directory
	c := fmt.Sprintf("%s%s%s", base, string(os.PathSeparator), "conf")
	return c
}

func CollectFiles(files *[]string, root string) error {
	return collectFiles(files, root, "")
}

// 调用时，root和prefix使用相同的值，则可以得到文件的绝对路径，但是prefix 不能以/结尾，需要处理
// root = / 时需要特殊处理
// todo change cwd wd 必须提供一个相对路径出来，方便将文件名作为对象名，但是又需要能够读取文件
//  set ENV? reset ENV?

// prefix 参数传入空字符串时，获取相对路径
func collectFiles(files *[]string, root, prefix string) error {

	fis, err := ioutil.ReadDir(root)
	if err != nil {
		return err
	}

	for _, v := range fis {
		file := ""

		if len(prefix) > 0 {
			file = fmt.Sprintf("%s%s%s", prefix, string(os.PathSeparator), v.Name())
		} else {
			file = v.Name()
		}

		if !v.IsDir() {
			*files = append(*files, file)
		} else {
			err := collectFiles(files, file, file)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
