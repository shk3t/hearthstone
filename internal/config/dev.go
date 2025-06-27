package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Config config

type config struct {
	PrintFrame          bool
	UnlimitedMana       bool
	RevealOpponentsHand bool
}

func Init() {
	godotenv.Load(".env")
	Config.PrintFrame = parseBool("PRINT_FRAME", false)
	Config.UnlimitedMana = parseBool("UNLIMITED_MANA", false)
	Config.UnlimitedMana = parseBool("UNLIMITED_MANA", false)
	Config.RevealOpponentsHand = parseBool("REVEAL_OPPONENTS_HAND", false)
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