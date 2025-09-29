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
		legacy.Neutral.RaidLeader,
		// legacy.Neutral.RiverCrocolisk,
		legacy.Neutral.ChillwindYeti,
		legacy.Neutral.ElvenArcher,
		legacy.Neutral.LootHoarder,
		legacy.Neutral.ColdlightOracle,
		// legacy.Mage.Frostbolt,
		// legacy.Mage.Fireball,
	)

	loop.StartGame(
		base.Heroes.Mage.Copy(),
		base.Heroes.Priest.Copy(),
		startingDeck.Copy(),
		startingDeck.Copy(),
	)
}