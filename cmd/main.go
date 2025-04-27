package main

import (
	"hearthstone/internal/cards"
	"hearthstone/internal/config"
	"hearthstone/internal/core"
	gamepkg "hearthstone/internal/game"
)

func main() {
	config.InitConfig()
	game := gamepkg.NewGame()

	game.TopPlayer.Hand[0] = &cards.AllCards.RiverCrocolisk
	game.TopPlayer.Hand[1] = &cards.AllCards.RiverCrocolisk
	game.TopPlayer.Hand[2] = &cards.AllCards.RiverCrocolisk
	game.TopPlayer.Hand[3] = &cards.AllCards.RiverCrocolisk
	game.TopPlayer.Hand[4] = &cards.AllCards.RiverCrocolisk
	game.TopPlayer.Hand[5] = &cards.AllCards.ChillwindYeti

	game.BotPlayer.Hand[3] = &cards.AllCards.Fireball

	core.StartGame(game)
}