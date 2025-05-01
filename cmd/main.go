package main

import (
	"hearthstone/internal/cards"
	"hearthstone/internal/core"
	"hearthstone/internal/game"
)

func main() {
	core.InitAll()
	defer core.DeinitAll()

	startingDeck := game.NewDeck(
		&cards.AllCards.RiverCrocolisk,
		&cards.AllCards.ChillwindYeti,
		&cards.AllCards.RiverCrocolisk,
		&cards.AllCards.ChillwindYeti,
	)

	core.StartGame(startingDeck, startingDeck)
}