package tui

import (
	"hearthstone/internal/game"
	"strings"
)

func HandleInput(g *game.Game) error {
	if !scanner.Scan() {
		return NewEndOfInputError()
	}

	input := scanner.Text()
	input = strings.ToLower(input)
	allArgs := strings.Fields(input)

	command, args := "", []string{}
	if len(allArgs) > 0 {
		command, args = allArgs[0], allArgs[1:]
	}

	for _, action := range actionList {
		if strings.HasPrefix(command, action.shortcut) || command == action.name {
			uiState.hint = action.Do(args, g)
			return nil
		}
	}

	uiState.hint = Actions.ShortHelp.Do(args, g)
	return nil
}
