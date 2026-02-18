package main

import (
	"hearthstone/internal/game"
	"hearthstone/internal/loop"
	"hearthstone/internal/sets/base"
	"hearthstone/internal/sets/legacy"
	"hearthstone/internal/tui"
	"hearthstone/pkg/helper"
)

func main() {
	loop.InitAll()
	defer loop.DeinitAll()

	startingDeck := game.NewDeck(
		// legacy.Mage.Frostbolt,
		// legacy.Mage.Fireball,
		// legacy.Neutral.QuestingAdventurer,
		legacy.Neutral.RaidLeader,
		legacy.Neutral.RaidLeader,
		// legacy.Neutral.ElvenArcher,
		legacy.Neutral.LootHoarder,
		legacy.Neutral.LootHoarder,
		legacy.Neutral.ColdlightOracle,
		legacy.Neutral.ColdlightOracle,
	)

	g := loop.StartGame(
		base.Heroes.Priest.Copy(),
		base.Heroes.Mage.Copy(),
		startingDeck.Copy(),
		startingDeck.Copy(),
		game.GameConfig{
			TableSize:           7,
			DisplayMethod:       game.DisplayMethods.TUI,
			FirstTurnSide:       game.BotSide,
			UnlimitedMana:       true,
			RevealOpponentsHand: false,
		},
		tui.NewGameIO(),
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
