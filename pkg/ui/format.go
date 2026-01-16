package ui

import "github.com/fatih/color"

type FormatFunc func(format string, a ...any) string

var BoldString = color.New(color.Bold).Sprintf
