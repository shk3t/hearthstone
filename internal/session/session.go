package session

import (
	"hearthstone/internal/game"
	"hearthstone/pkg/helpers"
)

type Session struct {
	*game.Game
	Hint         string // TODO extract (TUI related)
	TurnFinished bool
	Winner       game.Side
}

func NewGameSession(topHero, botHero *game.Hero, topDeck, botDeck game.Deck) *Session {
	return &Session{
		Game:         game.NewGame(topHero, botHero, topDeck, botDeck),
		Hint:         "",
		TurnFinished: true,
		Winner:       game.UnsetSide,
	}
}

func (s *Session) StartGame() {
	s.Game.StartGame()
}

func (s *Session) StartNextTurn() {
	s.TurnFinished = false
	s.Hint = ""
	errs := s.Game.StartNextTurn()
	if len(errs) > 0 {
		s.Hint = helpers.JoinErrors(errs, "\n")
	}
}

func (s *Session) Cleanup() {
	s.Game.Table.CleanupDeadMinions()
}

func (s *Session) CheckWinner() {
	for i := range game.SidesCount {
		side := game.Side(i)
		if !s.Game.Players[side].Hero.Alive {
			s.Winner = side.Opposite()
		}
	}
}

func (s *Session) HasWinner() bool {
	return s.Winner != game.UnsetSide
}