package main

import (
	"hearthstone/internal/cards"
	"hearthstone/internal/game"
	"hearthstone/internal/ui"
)

func main() {
	topPlayer := game.NewPlayer()
	botPlayer := game.NewPlayer()

	topPlayer.Hand[0] = &cards.AllCards.ChillwindYeti
	topPlayer.Hand[3] = &cards.AllCards.RiverCrocolisk

	botPlayer.Hand[3] = &cards.AllCards.Fireball

	table := &game.Table{}
	table.Top[0] = &cards.AllCards.ChillwindYeti

	ui.DisplayGame(topPlayer, botPlayer, table)
}