package main

import (
	"hearthstone/internal/game"
	"hearthstone/internal/loop"
	"hearthstone/internal/sets"
	"hearthstone/internal/sets/base"
	"hearthstone/internal/setup"
)

func main() {
	setup.InitAll()
	defer setup.DeinitAll()

	startingDeck := game.NewDeck(
		sets.LegacySet.RiverCrocolisk,
		sets.LegacySet.ChillwindYeti,
		sets.LegacySet.Frostbolt,
		sets.LegacySet.Fireball,
	)

	loop.StartGame(
		base.Heroes.Mage.Copy(),
		base.Heroes.Priest.Copy(),
		startingDeck.Copy(),
		startingDeck.Copy(),
	)
}