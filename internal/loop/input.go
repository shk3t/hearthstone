package loop

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
	allArgs := strings.Split(input, " ")
	command, args := allArgs[0], allArgs[1:]

	game.Help = ""
	switch {
	case strings.HasPrefix(command, "h") || command == "help":
		err = Actions.Help.Do(args, game)
	case strings.HasPrefix(command, "p") || command == "play":
		err = Actions.Play.Do(args, game)
	case strings.HasPrefix(command, "a") || command == "attack":
		err = Actions.Attack.Do(args, game)
	case strings.HasPrefix(command, "w") || command == "power":
		err = Actions.Power.Do(args, game)
	case strings.HasPrefix(command, "e") || command == "end":
		_ = Actions.End.Do(args, game)
	default:
		err = Actions.ShortHelp.Do(args, game)
	}

	if err != nil {
		game.Help = err.Error()
	}

	return false
}