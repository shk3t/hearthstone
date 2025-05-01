package core

import "errors"

func DoPlay(args []string, game *ActiveGame) error {
	if len(args) != 3 {
		return errors.New("Некорректные аргументы\n" + playDefaultHelp)
	}

	handIdx, err := parseIndexFromPosition(args[1])
	if err != nil {
		return errors.New("Некорретный 1 аргумент\n" + playDefaultHelp)
	}

	areaIdx, err := parseIndexFromPosition(args[1])
	if err != nil {
		return errors.New("Некорретный 2 аргумент\n" + playDefaultHelp)
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