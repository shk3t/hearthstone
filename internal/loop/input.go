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
	allArgs := strings.Fields(input)

	var command string
	var args []string
	if len(allArgs) > 0 {
		command, args = allArgs[0], allArgs[1:]
	}

	game.Help = ""

	err = Actions.ShortHelp.Do(args, game) // Display short help by default
	for _, action := range actionList {
		if strings.HasPrefix(command, action.shortcut) || command == action.name {
			err = action.Do(args, game)
			break
		}
	}

	if err != nil {
		game.Help = err.Error()
	}
	return false
}