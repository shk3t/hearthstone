package tui

import "regexp"

var actionCharRegexp = regexp.MustCompile("[ac-su-z]")
var multipleSpaceRegexp = regexp.MustCompile(" +")
var multipleBreakRegexp = regexp.MustCompile("\n+")
var actionArgumentRegexp = regexp.MustCompile("<(.*?)>")