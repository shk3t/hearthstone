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

func Feedback(errs ...error) {
	builder := strings.Builder{}
	for _, err := range errs {
		fmt.Fprintln(&builder, tuiError(err))
	}
	uiState.hint = strings.TrimSuffix(builder.String(), "\n")
}

var scanner = bufio.NewScanner(os.Stdin)

var uiState = struct {
	hint string
}{
	hint: "",
}