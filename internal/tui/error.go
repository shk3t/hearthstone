package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"hearthstone/internal/loop"
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

	case game.NoTargetSpecifiedError:
		out = "Цель не указана"

	case game.UnmatchedTargetNumberError:
		out = fmt.Sprintf(
			"Несоответствующее число целей.\nУказано: %d, требуется: %d",
			err.Specified, err.Required,
		)

	case game.UsedHeroPowerError:
		out = "Сила героя уже была использована в этом ходу"

	case game.UnavailableMinionAttackError:
		out = "Это существо сможет атаковать только в следующем ходу"

	case game.InfoError:
		switch card := err.Card.(type) {
		case game.Minion:
			return minionInfo(card)
		case game.Spell:
			return cardInfo(card.Card, color.MagentaString)
		default:
			panic("Invalid card type")
		}

	case loop.NeedActionError:
		return NewShortHelpError().Error()

	case InvalidArgumentsError, ShortHelpError, HelpError, InputPromptError:
		return err.Error()

	default:
		panic(errpkg.NewUnexpectedError(err))
	}

	return color.RedString(out)
}

type InvalidArgumentsError struct {
	action *tuiAction
}
type HelpError struct{}
type ShortHelpError struct{}
type InputPromptError struct {
	activeSide game.Side
}

func NewInvalidArgumentsError(action *tuiAction) InvalidArgumentsError {
	return InvalidArgumentsError{
		action: action,
	}
}
func NewShortHelpError() ShortHelpError {
	return ShortHelpError{}
}
func NewHelpError() HelpError {
	return HelpError{}
}
func NewInputPromptError(activeSide game.Side) InputPromptError {
	return InputPromptError{
		activeSide: activeSide,
	}
}

func (err InvalidArgumentsError) Error() string {
	builder := strings.Builder{}
	fmt.Fprintln(&builder, color.RedString("Некорректные аргументы"))
	if err.action != nil {
		fmt.Fprintln(&builder, err.action.info(true, false))
	}
	fmt.Fprint(&builder, positionsInfo())
	return builder.String()
}
func (err ShortHelpError) Error() string {
	builder := strings.Builder{}
	fmt.Fprint(&builder, color.RedString("Некорректное действие\n"))
	fmt.Fprint(&builder, color.YellowString("Доступные действия:\n"))
	for _, action := range actionList {
		fmt.Fprintln(&builder, action.info(false, true))
	}
	return strings.TrimSuffix(builder.String(), "\n")
}
func (err HelpError) Error() string {
	builder := strings.Builder{}
	fmt.Fprint(&builder,
		color.YellowString("Доступные действия:\n"),
	)
	for _, action := range actionList {
		fmt.Fprintln(&builder, action.info(false, false))
	}
	fmt.Fprint(&builder, positionsInfo())
	return builder.String()
}
func (err InputPromptError) Error() string {
	return getColorStringFunc(err.activeSide)("Выберите цели")
}

func positionsInfo() string {
	builder := strings.Builder{}
	fmt.Fprintf(&builder,
		"Чтобы указать героя в качестве цели, используйте %s\n",
		color.BlueString("0"),
	)
	fmt.Fprintf(&builder,
		"Чтобы указать сторону цели, используйте %s (верх) или %s (низ), например %s",
		color.BlueString("t"),
		color.BlueString("b"),
		color.BlueString("5b"),
	)
	return builder.String()
}
