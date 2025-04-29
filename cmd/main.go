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
	game.TopPlayer.Hand[1] = &cards.AllCards.ChillwindYeti
	game.TopPlayer.Hand[2] = &cards.AllCards.ChillwindYeti
	game.TopPlayer.Hand[3] = &cards.AllCards.ChillwindYeti
	game.TopPlayer.Hand[4] = &cards.AllCards.ChillwindYeti
	game.TopPlayer.Hand[5] = &cards.AllCards.ChillwindYeti
	game.TopPlayer.Hand[6] = &cards.AllCards.ChillwindYeti
	game.TopPlayer.Hand[7] = &cards.AllCards.ChillwindYeti

	game.TopPlayer.Deck[0] = &cards.AllCards.RiverCrocolisk
	game.TopPlayer.Deck[1] = &cards.AllCards.ChillwindYeti
	game.TopPlayer.Deck[2] = &cards.AllCards.ChillwindYeti
	game.TopPlayer.Deck[3] = &cards.AllCards.ChillwindYeti
	game.TopPlayer.Deck[4] = &cards.AllCards.ChillwindYeti
	game.TopPlayer.Deck[5] = &cards.AllCards.ChillwindYeti
	game.TopPlayer.Deck[6] = &cards.AllCards.ChillwindYeti
	game.TopPlayer.Deck[7] = &cards.AllCards.ChillwindYeti

	game.BotPlayer.Hand[0] = &cards.AllCards.Fireball

	core.StartGame(game)
}