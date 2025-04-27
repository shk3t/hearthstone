package core

import (
	"errors"
	"fmt"
	"hearthstone/internal/game"
	"strings"
)

func HandleInput(game *game.Game) {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}

	input = strings.ToLower(input)
	args := strings.Split(input, " ")

	game.InputError = nil
	switch {
	case strings.HasPrefix(args[0], "p"):
		break
	default:
		game.InputError = errors.New("Invalid action")
	}
}