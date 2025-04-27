package core

import (
	"bufio"
	"fmt"
	"hearthstone/internal/game"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func HandleInput(game *game.Game) {
	scanner.Scan()
	input := scanner.Text()

	input = strings.ToLower(input)
	args := strings.Split(input, " ")
	fmt.Println(args)

	game.InputHelp = ""
	switch {
	case strings.HasPrefix(args[0], "p"): // Play
		if len(args) != 3 {
			game.InputHelp = `Invalid arguments, use the following form:
play <hand position> <table position>
`
			break
		}

		handPos, err := strconv.Atoi(args[1])
		if err != nil {
			game.InputHelp = `Invalid 1 argument, use the following form:
play <hand position> <table position>
`
		}

		tablePos, err := strconv.Atoi(args[2])
		if err != nil {
			game.InputHelp = `Invalid 2 argument, use the following form:
play <hand position> <table position>
`
		}

		player := game.GetActivePlayer()
		player.PlayCard(handPos-1, tablePos-1)

	default:
		game.InputHelp = `Invalid action, actions available:
play (p) - play a card
`
	}
}