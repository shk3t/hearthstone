package game

import (
	"fmt"
	"hearthstone/internal/cards"
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
	BurnedCard cards.Playable
}
type InvalidTableAreaPositionError struct {
	position int
}
type FullTableAreaError struct{}
type EmptyDeckError struct {
	Fatigue int
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
func NewInvalidTableAreaPositionError(idx int) InvalidTableAreaPositionError {
	return InvalidTableAreaPositionError{position: idx + 1}
}
func NewFullTableAreaError() FullTableAreaError {
	return FullTableAreaError{}
}
func NewEmptyDeckError() EmptyDeckError {
	return EmptyDeckError{}
}

func (err CardPickError) Error() string {
	return fmt.Sprintf("Выбрана некорректная карта: %d", err.position)
}
func (err NotEnoughManaError) Error() string {
	return fmt.Sprintf("Недостаточно маны. Нужно: %d, имеется: %d", err.available, err.required)
}
func (err EmptyHandError) Error() string {
	return "Пустая рука"
}
func (err FullHandError) Error() string {
	if err.BurnedCard != nil {
		return fmt.Sprintf(
			"Полная рука. Последняя сожженная карта: \"%s\"",
			cards.ToCard(err.BurnedCard).Name,
		)
	}
	return "Полная рука"
}
func (err InvalidTableAreaPositionError) Error() string {
	return fmt.Sprintf("Некорректная позиция на столе: %d", err.position)
}
func (err FullTableAreaError) Error() string {
	return "Полный стол"
}
func (err EmptyDeckError) Error() string {
	if err.Fatigue != 0 {
		return fmt.Sprintf("Пустая колода. Потеря здоровья из-за усталости: %d", err.Fatigue)
	}
	return "Пустая колода"
}