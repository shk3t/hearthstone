package game

import (
	"fmt"
	"strings"
)

type CardPickError struct {
	position int
}
type NotEnoughManaError struct {
	available int
	required  int
}
type EmptyHandError struct{}
type FullHandError struct {
	BurnedCard Playable
}
type InvalidTableAreaPositionError struct {
	position int
	side     Side
}
type FullTableAreaError struct{}
type EmptyDeckError struct {
	Fatigue int
}
type UnmatchedEffectsAndTargetsError struct {
	spellName  string
	effectsLen int
	targetsLen int
}
type InvalidTargettingError struct {
	speicified int
	required   int
}

func NewCardPickError(idx int) CardPickError {
	return CardPickError{position: idx + 1}
}
func NewNotEnoughManaError(available, required int) NotEnoughManaError {
	return NotEnoughManaError{available, required}
}
func NewEmptyHandError() EmptyHandError {
	return EmptyHandError{}
}
func NewFullHandError() FullHandError {
	return FullHandError{}
}
func NewInvalidTableAreaPositionError(idx int, side Side) InvalidTableAreaPositionError {
	return InvalidTableAreaPositionError{
		position: idx + 1,
		side:     side,
	}
}
func NewFullTableAreaError() FullTableAreaError {
	return FullTableAreaError{}
}
func NewEmptyDeckError() EmptyDeckError {
	return EmptyDeckError{}
}
func NewUnmatchedEffectsAndTargetsError[T any](
	spell *Spell,
	targets []T,
) UnmatchedEffectsAndTargetsError {
	return UnmatchedEffectsAndTargetsError{
		spellName:  spell.Name,
		effectsLen: len(spell.TargetEffects),
		targetsLen: len(targets),
	}
}
func NewInvalidTargettingError(specified, required int) InvalidTargettingError {
	return InvalidTargettingError{
		speicified: specified,
		required:   required,
	}
}

func (err CardPickError) Error() string {
	return fmt.Sprintf("Выбрана некорректная карта: %d", err.position)
}
func (err NotEnoughManaError) Error() string {
	return fmt.Sprintf(
		"Недостаточно маны. Нужно: %d, имеется: %d",
		err.required,
		err.available,
	)
}
func (err EmptyHandError) Error() string {
	return "Пустая рука"
}
func (err FullHandError) Error() string {
	if err.BurnedCard != nil {
		return fmt.Sprintf(
			"Полная рука. Последняя сожженная карта: \"%s\"",
			ToCard(err.BurnedCard).Name,
		)
	}
	return "Полная рука"
}
func (err InvalidTableAreaPositionError) Error() string {
	if err.side == UnsetSide {
		return fmt.Sprintf("Некорректная позиция на столе: %d", err.position)
	}

	sideText := strings.ToLower(err.side.String())
	sideText = strings.Replace(sideText, "ий", "ей", 1)
	return fmt.Sprintf(
		"Некорректная позиция на %s части стола: %d",
		sideText,
		err.position,
	)
}
func (err FullTableAreaError) Error() string {
	return "Полный стол"
}
func (err EmptyDeckError) Error() string {
	if err.Fatigue != 0 {
		return fmt.Sprintf("Пустая колода.\nПотеря здоровья из-за усталости: %d", err.Fatigue)
	}
	return "Пустая колода"
}
func (err UnmatchedEffectsAndTargetsError) Error() string {
	return fmt.Sprintf(
		"Число эффектов и целей не соответствует для \"%s\".\nЭффектов: %d, целей: %d",
		err.spellName, err.effectsLen, err.targetsLen,
	)
}
func (err InvalidTargettingError) Error() string {
	return fmt.Sprintf(
		"Некорректный выбор цели.\nУказано целей: %d, требуется: %d",
		err.speicified, err.required,
	)
}