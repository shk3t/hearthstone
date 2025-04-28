package core

import (
	"bufio"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func HandleInput(game *ActiveGame) {
	var err error
	game.InputHelp = ""

	scanner.Scan()
	input := scanner.Text()

	input = strings.ToLower(input)
	args := strings.Split(input, " ")

	switch {
	case strings.HasPrefix(args[0], "p"):
		err = DoPlay(args, game)
	case strings.HasPrefix(args[0], "e"):
		DoEnd(game)
	default:
		game.InputHelp = actionsHelp
	}

	if err != nil {
		game.InputHelp = err.Error()
	}
}