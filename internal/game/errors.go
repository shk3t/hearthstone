package game

import errorpkg "hearthstone/pkg/errors"

// All game errors have to embed this structure.
// In this case they become instances of `error` type.
//
// `BaseError` `Error()` is not supposed to be used.
// Use your to string conversion function instead.
type BaseError struct{}

func (err BaseError) Error() string {
	panic(errorpkg.NewUnusableFeatureError())
}

type CardPickError struct {
	BaseError
	Position int
}
type NotEnoughManaError struct {
	BaseError
	Available int
	Required  int
}
type EmptyHandError struct {
	BaseError
}
type FullHandError struct {
	BaseError
	BurnedCard Playable
}
type InvalidTableAreaPositionError struct {
	BaseError
	Position int
	Side     Side
}
type FullTableAreaError struct {
	BaseError
}
type EmptyDeckError struct {
	BaseError
	Fatigue int
}
type UnmatchedTargetNumberError struct {
	BaseError
	Specified int
	Required  int
}
type UsedHeroPowerError struct {
	BaseError
}
type UnavailableMinionAttackError struct {
	BaseError
}

func NewCardPickError(idx int) CardPickError {
	return CardPickError{Position: idx + 1}
}
func NewNotEnoughManaError(available, required int) NotEnoughManaError {
	return NotEnoughManaError{Available: available, Required: required}
}
func NewEmptyHandError() EmptyHandError {
	return EmptyHandError{}
}
func NewFullHandError() FullHandError {
	return FullHandError{}
}
func NewInvalidTableAreaPositionError(idx int, side Side) InvalidTableAreaPositionError {
	return InvalidTableAreaPositionError{
		Position: idx + 1,
		Side:     side,
	}
}
func NewFullTableAreaError() FullTableAreaError {
	return FullTableAreaError{}
}
func NewEmptyDeckError() EmptyDeckError {
	return EmptyDeckError{}
}
func NewUnmatchedTargetNumberError(specified, required int) UnmatchedTargetNumberError {
	return UnmatchedTargetNumberError{
		Specified: specified,
		Required:  required,
	}
}
func NewUsedHeroPowerError() UsedHeroPowerError {
	return UsedHeroPowerError{}
}
func NewUnavailableMinionAttackError() UnavailableMinionAttackError {
	return UnavailableMinionAttackError{}
}