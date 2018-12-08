package flagscheck

import (
	"fmt"
	"github.com/kangaloo/cloudcli/display"
	"strings"
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

func NewConflictFlagErr(flags []string) error {
	return &conflictFlagErr{
		flags: flags,
	}
}

type conflictFlagErr struct {
	flags []string
}

func (e *conflictFlagErr) Error() string {
	var flags []string
	for _, flag := range e.flags {
		flags = append(flags, display.PrettyFlag(flag))
	}

	return fmt.Sprintf("conflict flags %s provided at the same time", strings.Join(flags, ", "))
}

func NewNotDefinedFlagErr(flag string) error {
	return &notDefinedFlagErr{
		flag: flag,
	}
}

type notDefinedFlagErr struct {
	flag string
}

func (e *notDefinedFlagErr) Error() string {
	return fmt.Sprintf("flag '%s' is flag checklist not defined", display.PrettyFlag(e.flag))
}

func NewAtLeastOneErr(flags []string) error {
	return &mustProvideOneErr{
		flags: flags,
	}
}

type mustProvideOneErr struct {
	flags []string
}

func (e *mustProvideOneErr) Error() string {
	var flags []string
	for _, flag := range e.flags {
		flags = append(flags, display.PrettyFlag(flag))
	}

	return fmt.Sprintf("at least one flag %s must be provided", strings.Join(flags, ", "))
}
