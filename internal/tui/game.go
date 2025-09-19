package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"strings"
)

const prompt = "> "

func gameString(g *game.Game) string {
	builder := strings.Builder{}

	fmt.Fprint(&builder, playerString(&g.Players[game.TopSide]))
	fmt.Fprint(&builder, tableString(&g.Table))
	fmt.Fprint(&builder, playerString(&g.Players[game.BotSide]))
	builder.WriteString("\n")

	if state.hint != "" {
		fmt.Fprintf(&builder, "%s\n\n", state.hint)
	}

	if winner := g.GetWinner(); winner != game.UnsetSide {
		fmt.Fprintf(
			&builder,
			"%s игрок одерживает ПОБЕДУ!\n",
			strings.ToUpper(winner.String()),
		)
	} else {
		fmt.Fprint(&builder, prompt)
	}

	return builder.String()
}