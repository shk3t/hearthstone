package tui

import (
	"fmt"
	"hearthstone/internal/game"
	errpkg "hearthstone/pkg/error"
	"strings"
)

func tuiError(err error) string {
	switch err := err.(type) {

	case nil:
		return ""

	case game.CardPickError:
		return fmt.Sprintf("Выбрана некорректная карта: %d", err.Position)

	case game.NotEnoughManaError:
		return fmt.Sprintf(
			"Недостаточно маны. Нужно: %d, имеется: %d",
			err.Required,
			err.Available,
		)

	case game.EmptyHandError:
		return "Пустая рука"

	case game.FullHandError:
		if err.BurnedCard != nil {
			return fmt.Sprintf(
				"Полная рука. Последняя сожженная карта: \"%s\"",
				game.ToCard(err.BurnedCard).Name,
			)
		}
		return "Полная рука"

	case game.InvalidTableAreaPositionError:
		if err.Side == game.UnsetSide {
			return fmt.Sprintf("Некорректная позиция на столе: %d", err.Position)
		}

		sideText := strings.ToLower(sideString(err.Side))
		sideText = strings.Replace(sideText, "ий", "ей", 1)
		return fmt.Sprintf(
			"Некорректная позиция на %s части стола: %d",
			sideText,
			err.Position,
		)

	case game.FullTableAreaError:
		return "Полный стол"

	case game.EmptyDeckError:
		if err.Fatigue != 0 {
			return fmt.Sprintf("Пустая колода.\nПотеря здоровья из-за усталости: %d", err.Fatigue)
		}
		return "Пустая колода"

	case game.UnmatchedTargetNumberError:
		return fmt.Sprintf(
			"Несоответствующее число целей.\nУказано: %d, требуется: %d",
			err.Specified, err.Required,
		)

	case game.UsedHeroPowerError:
		return "Сила героя уже была использована в этом ходу"

	case game.UnavailableMinionAttackError:
		return "Это существо сможет атаковать только в следующем ходу"

	default:
		panic(errpkg.NewUnexpectedError(err))
	}
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
	return strings.TrimSuffix(
		"Некорректные аргументы\n"+err.correctUsage,
		"\n",
	)
}
func (err EndOfInputError) Error() string {
	return "Конец ввода"
}