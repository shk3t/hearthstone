package tui

import (
	"fmt"
	"hearthstone/internal/game"
	errpkg "hearthstone/pkg/errors"
	"strings"

	"github.com/fatih/color"
)

func tuiError(err error) string {
	var out string

	switch err := err.(type) {

	case nil:
		out = ""

	case game.CardPickError:
		out = fmt.Sprintf("Выбрана некорректная карта: %d", err.Position)

	case game.NotEnoughManaError:
		out = fmt.Sprintf(
			"Недостаточно маны. Нужно: %d, имеется: %d",
			err.Required,
			err.Available,
		)

	case game.EmptyHandError:
		out = "Пустая рука"

	case game.FullHandError:
		if err.BurnedCard != nil {
			out = fmt.Sprintf(
				"Полная рука. Последняя сожженная карта: \"%s\"",
				game.ToCard(err.BurnedCard).Name,
			)
		}
		out = "Полная рука"

	case game.InvalidTableAreaPositionError:
		if err.Side == game.UnsetSide {
			out = fmt.Sprintf("Некорректная позиция на столе: %d", err.Position)
		}

		sideText := strings.ToLower(err.Side.String())
		sideText = strings.Replace(sideText, "ий", "ей", 1)
		out = fmt.Sprintf(
			"Некорректная позиция на %s части стола: %d",
			sideText,
			err.Position,
		)

	case game.FullTableAreaError:
		out = "Полный стол"

	case game.EmptyDeckError:
		if err.Fatigue != 0 {
			out = fmt.Sprintf("Пустая колода.\nПотеря здоровья из-за усталости: %d", err.Fatigue)
		}
		out = "Пустая колода"

	case game.UnmatchedTargetNumberError:
		out = fmt.Sprintf(
			"Несоответствующее число целей.\nУказано: %d, требуется: %d",
			err.Specified, err.Required,
		)

	case game.UsedHeroPowerError:
		out = "Сила героя уже была использована в этом ходу"

	case game.UnavailableMinionAttackError:
		out = "Это существо сможет атаковать только в следующем ходу"

	default:
		panic(errpkg.NewUnexpectedError(err))
	}

	return color.RedString(out)
}

type InvalidArgumentsError struct {
	correctUsage string
}
type EndOfInputError struct {
}

func (err InvalidArgumentsError) Set(correctUsage string) InvalidArgumentsError {
	err.correctUsage = correctUsage
	return err
}

func NewInvalidArgumentsError() InvalidArgumentsError {
	return InvalidArgumentsError{}
}
func NewEndOfInputError() EndOfInputError {
	return EndOfInputError{}
}

func (err InvalidArgumentsError) Error() string {
	return fmt.Sprintf(
		"%s:\n%s",
		color.RedString("Некорректные аргументы"),
		err.correctUsage,
	)
}
func (err EndOfInputError) Error() string {
	return color.RedString("Конец ввода")
}
