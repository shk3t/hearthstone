package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"hearthstone/pkg/ui"
	"strings"

	"github.com/fatih/color"
)

const prompt = "> "

func gameString(g *game.Game) string {
	builder := strings.Builder{}

	fmt.Fprint(&builder, playerString(&g.Players[game.TopSide]))
	fmt.Fprint(&builder, tableString(&g.Table, g.Turn))
	fmt.Fprint(&builder, playerString(&g.Players[game.BotSide]))
	fmt.Fprint(&builder, "\n\n")

	if state.hint != "" {
		fmt.Fprintf(&builder, "%s\n\n", state.hint)
	}

	if winner := g.GetWinner(); winner != game.UnsetSide {
		fmt.Fprintf(&builder,
			"%s %s\n",
			ui.BoldString(getColorStringFunc(winner)(
				strings.ToUpper(winner.String()),
			)),
			ui.BoldString(color.YellowString("игрок одерживает ПОБЕДУ!")),
		)
	} else {
		fmt.Fprint(
			&builder,
			ui.BoldString(getColorStringFunc(g.Turn)(prompt)),
		)
	}

	return builder.String()
}
