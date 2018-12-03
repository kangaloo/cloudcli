package file

import (
	"fmt"
	"os"
	"path/filepath"
)

func FileExist(file string) bool {
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

func SmartSize(s int64) string {

	if s < 1024 {
		return fmt.Sprintf("%dB", s)
	}

	if s < 1024*1024 {
		return fmt.Sprintf("%dKB", s/1024)
	}

	if s < 1024*1024*1024 {
		return fmt.Sprintf("%dMB", s/1024/1024)
	}

	if s < 1024*1024*1024*1024 {
		return fmt.Sprintf("%dGB", s/1024/1024/1024)
	}

	return fmt.Sprintf("%dTB", s/1024/1024/1024/1024)
}