package cloud

import (
	"fmt"
	"github.com/kangaloo/cloudcli/display"
)

func NewNecessaryFlagErr(flag string) error {
	return &necessaryFlagErr{flag: flag}
}

type necessaryFlagErr struct {
	flag string
}

func (e *necessaryFlagErr) Error() string {
	return fmt.Sprintf("missing necessary flag '%s'", display.PrettyFlag(e.flag))
}
