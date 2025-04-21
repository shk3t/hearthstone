package main

import (
	"fmt"
	"hearthstone/internal/game"
	s "strings"
)

func DisplayGame() {
	topPlayer := game.NewPlayer()
	botPlayer := game.NewPlayer()

	fmt.Println(topPlayer.Hero)

	fmt.Println(s.Repeat("=", 30))
	fmt.Println(s.Repeat("-", 30))
	fmt.Println(s.Repeat("=", 30))

	fmt.Println(botPlayer.Hero)
}

func main() {
	DisplayGame()
}