package game

import (
	"hearthstone/internal/config"
	"math/rand"
)

type Game struct {
	Players       [SidesCount]Player
	Table         Table
	Turn          Side
	TurnFinished  bool
	statusEffects map[*Character]PassiveEffect
	eventEffects  map[int]map[*Character]TriggerEffect
}

func NewGame(topHero, botHero *Hero, topDeck, botDeck Deck) *Game {
	game := &Game{
		Table:         *NewTable(),
		Turn:          UnsetSide,
		TurnFinished:  false,
		statusEffects: map[*Character]PassiveEffect{},
		eventEffects:  map[int]map[*Character]TriggerEffect{},
	}

	topHero.SetHealthToMax()
	botHero.SetHealthToMax()
	game.Players = [SidesCount]Player{
		TopSide: *newPlayer(TopSide, topHero, topDeck, game),
		BotSide: *newPlayer(BotSide, botHero, botDeck, game),
	}
	topHero.owner = &game.Players[0]
	botHero.owner = &game.Players[1]

	return game
}

func (g *Game) GetActivePlayer() *Player {
	return &g.Players[g.Turn]
}

func (g *Game) GetActiveArea() TableArea {
	return g.Table[g.Turn]
}

func (g *Game) StartGame() {
	turn := Side(config.Env.FirstTurnSide)
	if turn == UnsetSide {
		turn = Side(rand.Int() % 2)
	}
	firstPlayer, secondPlayer := g.Players[turn], g.Players[turn.Opposite()]

	firstPlayer.DrawCards(3)
	secondPlayer.DrawCards(4)
	secondPlayer.Hand.refill(BaseCards.TheCoin)

	g.Turn = turn.Opposite()
	g.StartNextTurn()
}

func (g *Game) StartNextTurn() []error {
	g.TurnFinished = false
	g.Turn = g.Turn.Opposite()

	activePlayer := g.GetActivePlayer()
	activePlayer.increaseMana()
	activePlayer.restoreMana()
	activePlayer.Hero.PowerIsUsed = false
	errs := activePlayer.DrawCards(1)

	activeArea := g.GetActiveArea()
	statuses := []*CharacterStatus{&activePlayer.Hero.Status}
	for _, minion := range activeArea.Minions {
		if minion != nil {
			statuses = append(statuses, &minion.Character.Status)
		}
	}
	for _, status := range statuses {
		status.SetSleep(false)
		status.Unfreeze()
	}

	return errs
}

func (g *Game) Cleanup() {
	for i := range SidesCount {
		side := Side(i)
		g.Table[side].cleanupDeadMinions()
	}
}

func (g *Game) GetWinner() Side {
	for i := range SidesCount {
		side := Side(i)
		if g.Players[side].Hero.Health <= 0 {
			return side.Opposite()
		}
	}
	return UnsetSide
}

func (g *Game) getCharacter(idx int, side Side) (*Character, error) {
	if idx == HeroIdx {
		return &g.Players[side].Hero.Character, nil
	} else {
		minion, err := g.Table[side].GetMinion(idx)
		if err != nil {
			return nil, err
		}
		return &minion.Character, nil
	}
}

func (g *Game) getApplicableStatusEffects(character *Character) []PassiveEffect {
	applicableEffects := []PassiveEffect{}
	for source, effect := range g.statusEffects {
		targets, _ := effect.Target(source, nil, nil)
		for _, target := range targets {
			if target == character {
				applicableEffects = append(applicableEffects, effect)
			}
		}
	}
	return applicableEffects
}
