package cloud

import "fmt"

func NewNecessaryFlagErr(flag string) error {
	return &necessaryFlagErr{flag: flag}
}

type necessaryFlagErr struct {
	flag string
}

func (e *necessaryFlagErr) Error() string {

	var flag string

	if len(e.flag) == 1 {
		flag = "-" + e.flag
	}

	if len(e.flag) > 1 {
		flag = "--" + e.flag
	}

	return fmt.Sprintf("missing necessary flag '%s'", flag)
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

func (e *conflictFlagErr) Error() string {
	return fmt.Sprintf("flag '%s' conflict with '%s'", e.flag, e.conflictFlag)
}

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
	return fmt.Sprintf("one flag must be chosen between '%s' and '%s'", e.flag, e.either)
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
	return fmt.Sprintf("flag '%s' is flag checklist not defined", e.flag)
}
