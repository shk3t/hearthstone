package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"strings"
)

var sideStrings = [game.SidesCount]string{"Верхний", "Нижний"}

func sideString(s game.Side) string {
	return sideStrings[s]
}

const prompt = "> "

func gameString(g *game.Game) string {
	builder := strings.Builder{}

	fmt.Fprint(&builder, playerString(&g.Players[game.TopSide]))
	fmt.Fprint(&builder, tableString(&g.Table))
	fmt.Fprint(&builder, playerString(&g.Players[game.BotSide]))

	if uiState.hint != "" {
		fmt.Fprintf(&builder, "%s\n\n", uiState.hint)
	}

	if winner := g.GetWinner(); winner != game.UnsetSide {
		fmt.Fprintf(
			&builder,
			"%s игрок одерживает ПОБЕДУ!\n",
			strings.ToUpper(sideString(winner)),
		)
	} else {
		fmt.Fprint(&builder, prompt)
	}

	return builder.String()
}