package flagscheck

import (
	"fmt"
	"github.com/kangaloo/cloudcli/display"
)

func NewEitherOrFlagErr(flag, either string) error {
	return &eitherOrFlagErr{
		flag:   flag,
		either: either,
	}
}

type eitherOrFlagErr struct {
	flag   string
	either string
}

func (e *eitherOrFlagErr) Error() string {
	return fmt.Sprintf("either '%s' or '%s' must be provided", display.PrettyFlag(e.flag), display.PrettyFlag(e.either))
}

func NewNecessaryFlagErr(flag string) error {
	return &necessaryFlagErr{flag: flag}
}

type necessaryFlagErr struct {
	flag string
}

func (e *necessaryFlagErr) Error() string {
	return fmt.Sprintf("missing necessary flag '%s'", display.PrettyFlag(e.flag))
}

func NewConflictFlagErr(flag, conflictFlag string) error {
	return &conflictFlagErr{
		flag:         flag,
		conflictFlag: conflictFlag,
	}
}

type conflictFlagErr struct {
	flag         string
	conflictFlag string
}

func NewNotDefinedFlagErr(flag string) error {
	return &notDefinedFlagErr{
		flag: flag,
	}
}

type notDefinedFlagErr struct {
	flag string
}

func (e notDefinedFlagErr) Error() string {
	return fmt.Sprintf("flag '%s' is flag checklist not defined", display.PrettyFlag(e.flag))
}
