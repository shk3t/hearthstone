package tui

import (
	"bufio"
	"fmt"
	"hearthstone/internal/game"
	"hearthstone/pkg/helper"
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

	if state.nextAction != nil {
		if actions.cancel.matches(command) {
			state.nextAction.rollback()
			state.hint, state.nextAction = helper.Capitalize(actions.cancel.description), nil
		} else {
			state.hint, state.nextAction = state.nextAction.wrappedDo(allArgs, g)
		}
		return nil
	}

	for _, action := range actionList {
		if action.matches(command) {
			state.hint, state.nextAction = action.wrappedDo(args, g)
			if state.nextAction != nil {
				state.hint = nextActionHint
			}
			return nil
		}
	}

	state.hint, _ = actions.shortHelp.wrappedDo(args, g)
	return nil
}

func Feedback(errs ...error) {
	builder := strings.Builder{}
	for _, err := range errs {
		fmt.Fprintln(&builder, tuiError(err))
	}
	state.hint = strings.TrimSuffix(builder.String(), "\n")
}

var scanner = bufio.NewScanner(os.Stdin)