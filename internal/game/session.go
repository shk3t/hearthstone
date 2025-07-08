package game

import (
	"hearthstone/pkg/helpers"
)

type Session struct {
	*Game
	Help         string
	TurnFinished bool
	Winner       Side
}

func NewGameSession(topHero, botHero *Hero, topDeck, botDeck Deck) *Session {
	return &Session{
		Game:         NewGame(topHero, botHero, topDeck, botDeck),
		Help:         "",
		TurnFinished: true,
		Winner:       UnsetSide,
	}
}

func (s *Session) StartGame() {
	s.Game.StartGame()
}

func (s *Session) StartNextTurn() {
	s.TurnFinished = false
	s.Help = ""
	errs := s.Game.StartNextTurn()
	if len(errs) > 0 {
		s.Help = helpers.JoinErrors(errs, "\n")
	}
}

func (s *Session) Cleanup() {
	s.Game.Table.CleanupDeadMinions()
}

func (s *Session) CheckWinner() {
	for i := range SidesCount {
		side := Side(i)
		if !s.Game.Players[side].Hero.Alive {
			s.Winner = side.Opposite()
		}
	}
}

func (s *Session) HasWinner() bool {
	return s.Winner != UnsetSide
}