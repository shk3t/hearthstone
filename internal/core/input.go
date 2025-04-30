package core

import (
	"bufio"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func handleInput(game *ActiveGame) {
	var err error

	if !scanner.Scan() {
		scanner = bufio.NewScanner(os.Stdin)
	}
	input := scanner.Text()

	input = strings.ToLower(input)
	args := strings.Split(input, " ")

	game.Help = ""
	switch {
	case strings.HasPrefix(args[0], "p"):
		err = DoPlay(args, game)
	case strings.HasPrefix(args[0], "e"):
		DoEnd(game)
	default:
		game.Help = actionsHelp
	}

	if err != nil {
		game.Help = err.Error()
	}
}