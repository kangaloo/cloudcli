package display

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
