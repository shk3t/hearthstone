package core

import (
	"errors"
)

func DoPlay(args []string, game *ActiveGame) error {
	if len(args) != 3 {
		return errors.New("Некорректные аргументы\n" + playUsageHelp)
	}

	handIdx, err := parseIndexFromPosition(args[1])
	if err != nil {
		return errors.New("Некорретный 1 аргумент\n" + playUsageHelp)
	}

	areaIdx, err := parseIndexFromPosition(args[1])
	if err != nil {
		return errors.New("Некорретный 2 аргумент\n" + playUsageHelp)
	}

	err = game.GetActivePlayer().PlayCard(handIdx, areaIdx)
	if err != nil {
		return err
	}
	return nil
}

func DoAttack(args []string, game *ActiveGame) error {
	if len(args) != 3 {
		return errors.New("Некорректные аргументы\n" + attackUsageHelp)
	}

	allyIdx, err := parseIndexFromPosition(args[1]) //TODO: parse hero
	if err != nil {
		return errors.New("Некорретный 1 аргумент\n" + attackUsageHelp)
	}

	enemyIdx, err := parseIndexFromPosition(args[1]) //TODO: parse hero
	if err != nil {
		return errors.New("Некорретный 2 аргумент\n" + attackUsageHelp)
	}

	err = game.GetActivePlayer().Attack(allyIdx, enemyIdx)
	if err != nil {
		return err
	}
	return nil
}

func DoEnd(game *ActiveGame) {
	game.TurnFinished = true
}