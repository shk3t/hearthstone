package game

import (
	"fmt"
	"hearthstone/internal/cards"
)

type CardPickError struct{}
type EmptyHandError struct{}
type FullHandError struct {
	BurnedCard cards.Playable
}
type FullTableAreaError struct{}
type EmptyDeckError struct {
	Fatigue int
}

func NewCardPickError() CardPickError {
	return CardPickError{}
}
func NewEmptyHandError() EmptyHandError {
	return EmptyHandError{}
}
func NewFullHandError() FullHandError {
	return FullHandError{}
}
func NewFullTableAreaError() FullTableAreaError {
	return FullTableAreaError{}
}
func NewEmptyDeckError() EmptyDeckError {
	return EmptyDeckError{}
}

func (err CardPickError) Error() string {
	return "Invalid card pick"
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
func (err FullTableAreaError) Error() string {
	return "Table is full"
}
func (err EmptyDeckError) Error() string {
	if err.Fatigue != 0 {
		return fmt.Sprintf("Deck is empty. Fatigue health loss: %d", err.Fatigue)
	}
	return "Deck is empty"
}