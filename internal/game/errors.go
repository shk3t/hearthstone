package game

import (
	"fmt"
	"hearthstone/internal/cards"
)

type CardPickError struct {
	position int
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
	return fmt.Sprintf("Invalid card pick: %d", err.position)
}
func (err EmptyHandError) Error() string {
	return "Hand is empty"
}
func (err FullHandError) Error() string {
	if err.BurnedCard != nil {
		return fmt.Sprintf(
			"Hand is full. Recent card was burned: \"%s\"",
			cards.ToCard(err.BurnedCard).Name,
		)
	}
	return "Hand is full"
}
func (err InvalidTableAreaPositionError) Error() string {
	return fmt.Sprintf("Invalid table position: %d", err.position)
}
func (err FullTableAreaError) Error() string {
	return "Table is full"
}
func (err EmptyDeckError) Error() string {
	if err.Fatigue != 0 {
		return fmt.Sprintf("Deck is empty. Fatigue health loss: %d", err.Fatigue)
	}
	return "Deck is empty"
}