package helper

import "unicode/utf8"

func BeForm(noun string) string {
	r, _ := utf8.DecodeLastRuneInString(noun)
	if r == 's' {
		return "are"
	}
	return "is"
}