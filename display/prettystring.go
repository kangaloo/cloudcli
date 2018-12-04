package display

import "fmt"

func PrettyFlag(flag string) string {

	if len(flag) == 1 {
		flag = "-" + flag
		return flag
	}

	if len(flag) > 1 {
		flag = "--" + flag
		return flag
	}

	return flag
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
