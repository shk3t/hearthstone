package tui

import (
	"bufio"
	"fmt"
	"hearthstone/internal/game"
	"hearthstone/pkg/ui"
	"os"
	"strings"
)

func Display(g *game.Game) {
	ui.UpdateFrame(gameString(g))
}

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

	if uiState.nextAction != nil {
		// TODO wrap with arg parsing (maybe wrap as an extra action?)
		err := uiState.nextAction.Do(args)

		if err == nil {
			uiState.nextAction.OnSuccess()
		} else {
			uiState.nextAction.OnFail()
			uiState.hint = tuiError(err)
		}

		uiState.nextAction = nil
		return nil
	}

	for _, action := range actionList {
		if strings.HasPrefix(command, action.shortcut) || command == action.name {
			uiState.hint, uiState.nextAction = action.Do(args, g)
			// TODO: set hint for nextAction somehow
			return nil
		}
	}

	uiState.hint, _ = Actions.ShortHelp.Do(args, g)
	return nil
}

func Feedback(errs ...error) {
	builder := strings.Builder{}
	for _, err := range errs {
		fmt.Fprintln(&builder, tuiError(err))
	}
	uiState.hint = strings.TrimSuffix(builder.String(), "\n")
}

var uiState = struct {
	hint       string
	nextAction *game.NextAction
}{
	hint: "",
}

var scanner = bufio.NewScanner(os.Stdin)