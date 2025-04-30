package core

import "errors"

func DoPlay(args []string, game *ActiveGame) error {
	if len(args) != 3 {
		return errors.New("Invalid arguments." + playDefaultHelp)
	}

	handIdx, err := parseIndexFromPosition(args[1])
	if err != nil {
		return errors.New("Invalid 1 argument." + playDefaultHelp)
	}

	areaIdx, err := parseIndexFromPosition(args[1])
	if err != nil {
		return errors.New("Invalid 2 argument." + playDefaultHelp)
	}

	activePlayer := game.GetActivePlayer()
	err = activePlayer.PlayCard(handIdx, areaIdx)
	if err != nil {
		return err
	}

	return nil
}

func DoEnd(game *ActiveGame) {
	game.TurnFinished = true
}