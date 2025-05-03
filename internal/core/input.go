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
	case strings.HasPrefix(args[0], "h") || args[0] == "help":
		game.Help = fullHelp
	case strings.HasPrefix(args[0], "p") || args[0] == "play":
		err = DoPlay(args, game)
	case strings.HasPrefix(args[0], "a") || args[0] == "attack":
		err = DoAttack(args, game)
	case strings.HasPrefix(args[0], "w") || args[0] == "power":
		err = DoUseHeroPower()
	case strings.HasPrefix(args[0], "e") || args[0] == "end":
		DoEnd(game)
	default:
		game.Help = availableActions
	}

	if err != nil {
		game.Help = err.Error()
	}

	return false
}