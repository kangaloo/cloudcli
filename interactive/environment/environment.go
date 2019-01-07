package environment

import "fmt"

// 交互模式中的Env即非交互模式中的全局参数
type Env struct {
	Key   string
	Flag  string
	Value string
}

func NewEnvs(keys ...string) []Env {
	var envs []Env
	for _, key := range keys {
		var flag string

		if key == "" {
			continue
		}

		if len(key) == 1 {
			flag = "-" + key
		}

		if len(key) > 1 {
			flag = "--" + key
		}

		envs = append(envs, Env{Key: key, Flag: flag})
	}

	return envs
}

func (e *Env) Display() {
	fmt.Printf("%s = %s\n", e.Key, e.Value)
}
