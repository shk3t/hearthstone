package main

import (
	"hearthstone/internal/game"
	"hearthstone/internal/loop"
	"hearthstone/internal/sets/base"
	"hearthstone/internal/sets/legacy"
	"hearthstone/internal/setup"
)

func main() {
	setup.InitAll()
	defer setup.DeinitAll()

	startingDeck := game.NewDeck(
		legacy.Neutral.RiverCrocolisk,
		legacy.Neutral.ChillwindYeti,
		legacy.Mage.Frostbolt,
		legacy.Mage.Fireball,
	)

	loop.StartGame(
		base.Heroes.Mage.Copy(),
		base.Heroes.Priest.Copy(),
		startingDeck.Copy(),
		startingDeck.Copy(),
	)
}