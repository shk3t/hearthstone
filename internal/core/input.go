package core

import (
	"bufio"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func handleInput(game *ActiveGame) (exit bool) {
	var err error

	if !scanner.Scan() {
		return true
	}
	input := scanner.Text()

	input = strings.ToLower(input)
	args := strings.Split(input, " ")

	game.Help = ""
	switch {
	case strings.HasPrefix(args[0], "p"):
		err = DoPlay(args, game)
	case strings.HasPrefix(args[0], "a"):
		err = DoAttack(args, game)
	case strings.HasPrefix(args[0], "e"):
		DoEnd(game)
	case strings.HasPrefix(args[0], "h"):
		game.Help = fullHelp
	default:
		game.Help = availableActions
	}

	if err != nil {
		game.Help = err.Error()
	}

	return false
}