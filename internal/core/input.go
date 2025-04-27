package core

import (
	"bufio"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func HandleInput(game *ActiveGame) {
	scanner.Scan()
	input := scanner.Text()

	input = strings.ToLower(input)
	args := strings.Split(input, " ")

	game.InputHelp = ""
	switch {
	case strings.HasPrefix(args[0], "p"):
		err := DoPlay(args, game)
		if err != nil {
			game.InputHelp = err.Error()
		}
	default:
		game.InputHelp = `Invalid action, actions available:
play (p) - play a card
`
	}
}