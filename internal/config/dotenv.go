package config

import (
	"hearthstone/pkg/sugar"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/joho/godotenv"
)

var Env envFields

var DisplayMethods = struct {
	Tui string
}{Tui: "TUI"}

func LoadEnv() {
	godotenv.Load(".env")
	Env.DisplayMethod = strings.ToUpper(parseString("DISPLAY_METHOD", "TUI"))
	Env.UnlimitedMana = parseBool("UNLIMITED_MANA", false)
	Env.RevealOpponentsHand = parseBool("REVEAL_OPPONENTS_HAND", false)

	firstTurnSide := os.Getenv("FIRST_TURN_SIDE")
	Env.FirstTurnSide = -1
	if len(firstTurnSide) > 0 {
		switch unicode.ToLower([]rune(firstTurnSide)[0]) {
		case 't':
			Env.FirstTurnSide = 0
		case 'b':
			Env.FirstTurnSide = 1
		}
	}
}

type envFields struct {
	DisplayMethod       string
	FirstTurnSide       int
	UnlimitedMana       bool
	RevealOpponentsHand bool
}

func parseBool(variable string, defaultValue bool) bool {
	if os.Getenv(variable) == "" {
		return defaultValue
	}
	value, error := strconv.ParseBool(os.Getenv(variable))
	if error != nil {
		return defaultValue
	}
	return value
}

func parseString(variable string, defaultValue string) string {
	return sugar.If(os.Getenv(variable) != "", os.Getenv(variable), defaultValue)
}
