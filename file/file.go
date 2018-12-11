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

// prefix 是递归用的，调用的时候必须传入空字符串 ""
func CollectFiles(files *[]string, root, prefix string) error {

	fis, err := ioutil.ReadDir(root)
	if err != nil {
		return err
	}

	for _, v := range fis {
		file := ""

		if len(prefix) > 0 {
			file = fmt.Sprintf("%s%s%s", root, string(os.PathSeparator), v.Name())
		} else {
			file = v.Name()
		}

		if !v.IsDir() {
			*files = append(*files, file)
		} else {
			err := CollectFiles(files, file, file)
			if err != nil {
				return nil
			}
		}
	}

	return nil
}
