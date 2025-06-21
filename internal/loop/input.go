package loop

import (
	"bufio"
	"hearthstone/pkg/log"
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

	err = Actions.ShortHelp.Do(args, game) // Display short help by default
	for _, action := range actionList {
		log.DLog(action)
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