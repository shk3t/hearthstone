package core

import "strconv"

func DoPlay(game *ActiveGame, args []string) {
	if len(args) != 3 {
		game.InputHelp = `Invalid arguments, use the following form:
play <hand position> <table position>`
		return
	}

	handPos, err := strconv.Atoi(args[1])
	if err != nil {
		game.InputHelp = `Invalid 1 argument, use the following form:
play <hand position> <table position>`
		return
	}

	tablePos, err := strconv.Atoi(args[2])
	if err != nil {
		game.InputHelp = `Invalid 2 argument, use the following form:
play <hand position> <table position>`
		return
	}

	player := game.GetActivePlayer()
	player.PlayCard(handPos-1, tablePos-1)

}