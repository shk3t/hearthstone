package core

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func HandleInput(game *ActiveGame) {
	scanner.Scan()
	input := scanner.Text()

	input = strings.ToLower(input)
	args := strings.Split(input, " ")
	fmt.Println(args)

	game.InputHelp = ""
	switch {
	case strings.HasPrefix(args[0], "p"):
		DoPlay(game, args)
	default:
		game.InputHelp = `Invalid action, actions available:
play (p) - play a card
`
	}
}