package ui

import (
	"fmt"
	"hearthstone/internal/game"
)

func DisplayGame(topPlayer, botPlayer *game.Player, table *game.Table) {
	fmt.Println(&topPlayer.Hero)
	fmt.Print(&topPlayer.Hand)
	fmt.Print(table)
	fmt.Print(&botPlayer.Hand)
	fmt.Println(&botPlayer.Hero)
}