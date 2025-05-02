package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Config config

type config struct {
	PrintFrame bool
	FreeMana   bool
}

func Init() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file")
	}

	Config.PrintFrame = parseBool("PRINT_FRAME", false)
	Config.FreeMana = parseBool("FREE_MANA", false)
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