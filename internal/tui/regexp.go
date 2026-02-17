package tui

import "regexp"

var multipleSpaceRegexp = regexp.MustCompile(" +")
var multipleBreakRegexp = regexp.MustCompile("\n+")
var actionArgumentRegexp = regexp.MustCompile("<(.*?)>")