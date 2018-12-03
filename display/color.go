package display

import "github.com/fatih/color"

var (
	Red     = color.New(color.FgRed).SprintfFunc()
	Green   = color.New(color.FgGreen).SprintfFunc()
	Yellow  = color.New(color.FgYellow).SprintfFunc()
	HiBlack = color.New(color.FgHiBlack).SprintfFunc()
)
