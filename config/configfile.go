package config

import (
	"errors"
	"fmt"
	"github.com/kangaloo/cloudcli/file"
	"os"
)

var defaultFile = "cloudcli.yaml"

var (
	//NotExistErr = errors.New("config file not exist")
	//NotFileErr  = errors.New("config file is not a file")
	ZeroLenErr = errors.New("config file path is zero length")
	NotSetErr  = errors.New("global config file is not set")
)

func DefaultFile() string {
	path := file.DefaultConfDir()
	conf := fmt.Sprintf("%s%s%s", path, string(os.PathSeparator), defaultFile)
	return conf
}
