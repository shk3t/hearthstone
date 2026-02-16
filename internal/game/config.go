package game

type GameConfig struct {
	TableSize           int
	DisplayMethod       string
	FirstTurnSide       Side
	UnlimitedMana       bool
	RevealOpponentsHand bool
}

var DisplayMethods = struct {
	TUI string
}{TUI: "TUI"}

func DefaultGameConfig() GameConfig {
	return GameConfig{
		TableSize:           7,
		DisplayMethod:       DisplayMethods.TUI,
		FirstTurnSide:       BotSide,
		UnlimitedMana:       false,
		RevealOpponentsHand: false,
	}
}
