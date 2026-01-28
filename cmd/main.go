package main

import (
	"hearthstone/internal/game"
	"hearthstone/internal/loop"
	"hearthstone/internal/sets/base"
	"hearthstone/internal/sets/legacy"
	"hearthstone/internal/setup"
	"hearthstone/pkg/helper"
)

func main() {
	setup.InitAll()
	defer setup.DeinitAll()

	startingDeck := game.NewDeck(
		legacy.Neutral.QuestingAdventurer,
		legacy.Neutral.RaidLeader,
		legacy.Neutral.ElvenArcher,
		legacy.Neutral.LootHoarder,
		legacy.Neutral.ColdlightOracle,
	)

	g := loop.StartGame(
		base.Heroes.Mage.Copy(),
		base.Heroes.Priest.Copy(),
		startingDeck.Copy(),
		startingDeck.Copy(),
	)

	_ = g
	// topPlayer := g.Players[game.TopSide]
	// botPlayer := g.Players[game.BotSide]
	//
	// topPlayer.PlayCard(0, 0, nil, nil)
	// topPlayer.PlayCard(0, 0, nil, nil)
	// topPlayer.PlayCard(0, 0, nil, nil)
	// g.StartNextTurn()
	// botPlayer.PlayCard(0, 0, nil, nil)
	// botPlayer.PlayCard(0, 0, nil, nil)
	// botPlayer.PlayCard(0, 0, nil, nil)
	// g.StartNextTurn()

	helper.WaitForever()
}