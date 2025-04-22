package ui

import (
	"fmt"
	"hearthstone/internal/game"
	"strings"
)

func DisplayGame(topPlayer, botPlayer *game.Player, table *game.Table) {
	fmt.Println(&topPlayer.Hero)

	fmt.Print(&topPlayer.Hand)

	fmt.Println(strings.Repeat("=", 30))
	fmt.Println(&table.Top)
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println(&table.Bot)
	fmt.Println(strings.Repeat("=", 30))

	fmt.Print(&botPlayer.Hand)

	fmt.Println(&botPlayer.Hero)
}