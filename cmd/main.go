package main

import (
	"hearthstone/internal/sets"
	"hearthstone/internal/loop"
	"hearthstone/internal/game"
)

func main() {
	loop.InitAll()
	defer loop.DeinitAll()

	startingDeck := game.NewDeck(
		sets.LegacySet.RiverCrocolisk.Copy(),
		sets.LegacySet.RiverCrocolisk.Copy(),
		sets.LegacySet.RiverCrocolisk.Copy(),
		sets.LegacySet.RiverCrocolisk.Copy(),
	)

	loop.StartGame(startingDeck, startingDeck)
}