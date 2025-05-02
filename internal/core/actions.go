package core

import (
	"fmt"
)

func DoPlay(args []string, game *ActiveGame) error {
	if len(args) != 3 {
		return fmt.Errorf("Некорректные аргументы\n%s", actionsHelp.play.usage())
	}

	handIdx, err := parseIndexFromPosition(args[1]) // TEST
	if err != nil {
		return fmt.Errorf("Некорретный 1 аргумент\n%s", actionsHelp.play.usage())
	}

	areaIdx, err := parseIndexFromPosition(args[1]) // TEST
	if err != nil {
		return fmt.Errorf("Некорретный 2 аргумент\n%s", actionsHelp.play.usage())
	}

	err = game.GetActivePlayer().PlayCard(handIdx, areaIdx)
	if err != nil {
		return err
	}
	return nil
}

func DoAttack(args []string, game *ActiveGame) error {
	if len(args) != 3 {
		return fmt.Errorf("Некорректные аргументы\n%s", actionsHelp.attack.usage())
	}

	allyIdx, err := parseIndexFromPosition(args[1]) //TEST
	if err != nil {
		return fmt.Errorf("Некорретный 1 аргумент\n%s", actionsHelp.attack.usage())
	}

	enemyIdx, err := parseIndexFromPosition(args[2]) //TEST
	if err != nil {
		return fmt.Errorf("Некорретный 2 аргумент\n%s", actionsHelp.attack.usage())
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