package core

import (
	"errors"
	"strconv"
)

var playDefaultHelp = `
Use the following form:
play <hand position> <table position>`

func DoPlay(args []string, game *ActiveGame) error {
	if len(args) != 3 {
		return errors.New("Invalid arguments." + playDefaultHelp)
	}

	handPos, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.New("Invalid 1 argument." + playDefaultHelp)
	}

	tablePos, err := strconv.Atoi(args[2])
	if err != nil {
		return errors.New("Invalid 2 argument." + playDefaultHelp)
	}

	activePlayer := game.GetActivePlayer()
	err = activePlayer.PlayCard(handPos-1, tablePos-1)
	if err != nil {
		return err
	}

	return nil
}