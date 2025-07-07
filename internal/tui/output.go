package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"hearthstone/internal/setup"
	"hearthstone/pkg/sugar"
	"hearthstone/pkg/ui"
	"strings"
)

var DisplayFrame func(string)

// TODO: move to `setup` package
func InitDisplayFrame() {
	DisplayFrame = sugar.If(setup.Env.PrintFrame, ui.PrintFrame, ui.UpdateFrame)
}

func Display(gs *game.GameSession) {
	builder := strings.Builder{}
	fmt.Fprintln(&builder, gs.Game)

	if gs.Help != "" {
		fmt.Fprintf(&builder, "%s\n\n", gs.Help)
	}

	if gs.Winner != game.UnsetSide {
		fmt.Fprintf(
			&builder,
			"%s игрок одерживает ПОБЕДУ!\n",
			strings.ToUpper(gs.Winner.String()),
		)
	} else {
		fmt.Fprint(&builder, prompt)
	}

	DisplayFrame(builder.String())
}

const prompt = "> "