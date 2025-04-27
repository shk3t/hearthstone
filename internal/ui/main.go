package ui

import (
	"fmt"
	"hearthstone/internal/config"
	"hearthstone/internal/game"
)

func DisplayFrame(topPlayer, botPlayer *game.Player, table *game.Table) {
	if !config.Config.Debug {
		clearDisplay()
	}

	fmt.Print(&topPlayer.Hand)
	fmt.Println(&topPlayer.Hero)
	fmt.Print(table)
	fmt.Println(&botPlayer.Hero)
	fmt.Print(&botPlayer.Hand)

	if config.Config.Debug {
		displayFrameSeparator()
	}
}

func clearDisplay() {
	fmt.Print("\033[2J\033[H")
}