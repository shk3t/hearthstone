package main

import (
	"hearthstone/internal/game"
	"hearthstone/internal/loop"
	"hearthstone/internal/sets"
	"hearthstone/internal/sets/base"
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

	loop.StartGame(
		base.Heroes.Mage.Copy(),
		base.Heroes.Priest.Copy(),
		startingDeck,
		startingDeck,
	)
}