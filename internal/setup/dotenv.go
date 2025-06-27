package setup

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Env envFields

func Init() {
	godotenv.Load(".env")
	Env.PrintFrame = parseBool("PRINT_FRAME", false)
	Env.UnlimitedMana = parseBool("UNLIMITED_MANA", false)
	Env.UnlimitedMana = parseBool("UNLIMITED_MANA", false)
	Env.RevealOpponentsHand = parseBool("REVEAL_OPPONENTS_HAND", false)
}

type envFields struct {
	PrintFrame          bool
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