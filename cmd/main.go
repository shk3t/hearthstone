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
		cards.AllCards.RiverCrocolisk.Copy(),
		cards.AllCards.RiverCrocolisk.Copy(),
		cards.AllCards.RiverCrocolisk.Copy(),
		cards.AllCards.RiverCrocolisk.Copy(),
	)

	core.StartGame(startingDeck, startingDeck)
}