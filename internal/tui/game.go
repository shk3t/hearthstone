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

func sessionString(s *game.Session) string {
	builder := strings.Builder{}
	fmt.Fprintln(&builder, gameString(s.Game))

	if s.Help != "" {
		fmt.Fprintf(&builder, "%s\n\n", s.Help)
	}

	if s.Winner != game.UnsetSide {
		fmt.Fprintf(
			&builder,
			"%s игрок одерживает ПОБЕДУ!\n",
			strings.ToUpper(sideString(s.Winner)),
		)
	} else {
		fmt.Fprint(&builder, prompt)
	}

	return builder.String()
}

func gameString(g *game.Game) string {
	builder := strings.Builder{}
	fmt.Fprint(&builder, playerString(&g.Players[game.TopSide]))
	fmt.Fprint(&builder, tableString(&g.Table))
	fmt.Fprint(&builder, playerString(&g.Players[game.BotSide]))
	return builder.String()
}