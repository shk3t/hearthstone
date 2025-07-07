package game

import (
	"hearthstone/pkg/helpers"
)

type GameSession struct {
	*Game
	Help         string
	TurnFinished bool
	Winner       Side
}

func NewGameSession(topHero, botHero *Hero, topDeck, botDeck Deck) *GameSession {
	return &GameSession{
		Game:         NewGame(topHero, botHero, topDeck, botDeck),
		Help:         "",
		TurnFinished: true,
		Winner:       UnsetSide,
	}
}

func (gs *GameSession) StartGame() {
	gs.Game.StartGame()
}

func (gs *GameSession) StartNextTurn() {
	gs.TurnFinished = false
	gs.Help = ""
	errs := gs.Game.StartNextTurn()
	if len(errs) > 0 {
		gs.Help = helpers.JoinErrors(errs, "\n")
	}
}

func (gs *GameSession) Cleanup() {
	gs.Game.Table.CleanupDeadMinions()
}

func (gs *GameSession) CheckWinner() {
	for i := range SidesCount {
		side := Side(i)
		if !gs.Game.Players[side].Hero.Alive {
			gs.Winner = side.Opposite()
		}
	}
}

func (gs *GameSession) HasWinner() bool {
	return gs.Winner != UnsetSide
}