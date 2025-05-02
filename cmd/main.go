package main

import (
	"hearthstone/internal/cards/sets"
	"hearthstone/internal/core"
	"hearthstone/internal/game"
)

func main() {
	core.InitAll()
	defer core.DeinitAll()

	startingDeck := game.NewDeck(
		sets.LegacySet.RiverCrocolisk.Copy(),
		sets.LegacySet.RiverCrocolisk.Copy(),
		sets.LegacySet.RiverCrocolisk.Copy(),
		sets.LegacySet.RiverCrocolisk.Copy(),
	)

	core.StartGame(startingDeck, startingDeck)
}