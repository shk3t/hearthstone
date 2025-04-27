package ui

import "fmt"

func UpdateFrame(content string) {
	fmt.Print("\033[2J\033[H")
	fmt.Print(content)
}

func PrintFrame(content string) {
	fmt.Print("\n\n\n")
	fmt.Print(content)
}