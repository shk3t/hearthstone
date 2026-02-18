package game

import (
	"hearthstone/pkg/sugar"
	"math/rand"
)

type Game struct {
	Players       [SidesCount]Player
	Table         Table
	Turn          Side
	TurnFinished  bool
	eventEffects  map[int]map[*Character]Effect
	Config        GameConfig
	inputPosition func(n int) (idxes []int, sides []Side)
}

type PositionInputFunc func(g *Game, n int) (idxes []int, sides []Side)

func NewGame(
	topHero, botHero *Hero,
	topDeck, botDeck Deck,
	config GameConfig,
	positionInputFunc PositionInputFunc,
) *Game {
	var game *Game
	game = &Game{
		Table:        *newTable(config.TableSize),
		Turn:         UnsetSide,
		TurnFinished: false,
		eventEffects: map[int]map[*Character]Effect{},
		Config:       config,
		inputPosition: func(n int) (idxes []int, sides []Side) {
			return positionInputFunc(game, n)
		},
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
	turn := sugar.If(
		g.Config.FirstTurnSide != UnsetSide,
		g.Config.FirstTurnSide,
		Side(rand.Int()%2),
	)
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
	activePlayer.Hero.Power.IsUsed = false
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

func (g *Game) AttachEffect(effect Effect, source *Character) error {
	event := effect.Event
	if event.getPrimaryEvent != nil {
		event = event.getPrimaryEvent(source.owner)
	}

	characterEffects, ok := g.eventEffects[event.id]
	if !ok {
		characterEffects = map[*Character]Effect{}
		g.eventEffects[event.id] = characterEffects
	}
	characterEffects[source] = effect

	return nil
}

func (g *Game) DetachEffect(effect Effect, source *Character) error {
	event := effect.Event
	if event.getPrimaryEvent != nil {
		event = event.getPrimaryEvent(source.owner)
	}

	characterEffects, ok := g.eventEffects[event.id]
	if ok {
		delete(characterEffects, source)
	}

	return nil
}

func (g *Game) TriggerSelfEffect(event event, source *Character) error {
	for _, effect := range source.Effects {
		if effect.Event.id == event.id {
			return effect.Apply(source, nil, nil)
		}
	}
	return nil
}

func (g *Game) TriggerAllEffects(event event) {
	characterEffects := g.eventEffects[event.id]
	for character, effect := range characterEffects {
		effect.Apply(character, nil, nil)
	}
}

func (g *Game) TriggerAllEffectsOnlyFor(event event, target *Character) {
	characterEffects := g.eventEffects[event.id]
	for source, effect := range characterEffects {
		effect.ApplyOn(source, target)
	}
}

func (g *Game) GetInfo(idx int, side Side) error {
	player := g.GetActivePlayer()

	if idx == HeroIdx {
		if side == player.Side.Opposite() {
			return NewInfoError(player.GetOpponent().Hero.Power)
		}
		return NewInfoError(player.Hero.Power)
	}

	if side == UnsetSide {
		card, err := player.Hand.Get(idx)
		if err != nil {
			return err
		}
		return NewInfoError(card)
	}

	minion, err := g.Table[side].GetMinion(idx)
	if err != nil {
		return err
	}
	return NewInfoError(*minion)
}
