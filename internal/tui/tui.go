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

	for _, action := range actionList {
		if strings.HasPrefix(command, action.shortcut) || command == action.name {
			uiState.hint = action.Do(args, g)
			return nil
		}
	}

	uiState.hint = Actions.ShortHelp.Do(args, g)
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
	hint string
}{
	hint: "",
}

var scanner = bufio.NewScanner(os.Stdin)