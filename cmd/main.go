package main

import (
	"hearthstone/internal/cards"
	"hearthstone/internal/core"
	gamepkg "hearthstone/internal/game"
)

func main() {
	core.InitAll()
	defer core.DeinitAll()

	game := gamepkg.NewGame()

	game.TopPlayer.Hand[0] = &cards.AllCards.RiverCrocolisk
	game.TopPlayer.Hand[5] = &cards.AllCards.ChillwindYeti

	game.BotPlayer.Hand[3] = &cards.AllCards.Fireball

	core.StartGame(game)
}