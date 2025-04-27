package main

import (
	"hearthstone/internal/cards"
	"hearthstone/internal/config"
	"hearthstone/internal/game"
	"hearthstone/internal/ui"
)

func main() {
	config.InitConfig()
	table := game.NewTable()

	topPlayer := game.NewPlayer(table)
	botPlayer := game.NewPlayer(table)

	topPlayer.Hand[0] = &cards.AllCards.RiverCrocolisk
	topPlayer.Hand[1] = &cards.AllCards.RiverCrocolisk
	topPlayer.Hand[2] = &cards.AllCards.RiverCrocolisk
	topPlayer.Hand[3] = &cards.AllCards.RiverCrocolisk
	topPlayer.Hand[4] = &cards.AllCards.RiverCrocolisk
	topPlayer.Hand[5] = &cards.AllCards.ChillwindYeti

	botPlayer.Hand[3] = &cards.AllCards.Fireball

	ui.DisplayGame(topPlayer, botPlayer, table)

	topPlayer.PlayCard(0, 3)
	topPlayer.PlayCard(1, 3)
	topPlayer.PlayCard(2, 3)
	ui.DisplayGame(topPlayer, botPlayer, table)

	topPlayer.PlayCard(2, 2)
	ui.DisplayGame(topPlayer, botPlayer, table)
}