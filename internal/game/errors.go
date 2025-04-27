package game

import "fmt"

type CardPickError struct {
	err error
}
type EmptyHandError struct {
	err error
}
type FullTableAreaError struct {
	err error
}

func (err CardPickError) Error() string {
	return err.err.Error()
}
func (err EmptyHandError) Error() string {
	return "Hand is empty"
}
func (err FullTableAreaError) Error() string {
	return "Table is full"
}

func NewCardPickError(err error) CardPickError {
	return CardPickError{
		err: fmt.Errorf("Invalid card was picked\n%w", err),
	}
}
func NewEmptyHandError() EmptyHandError {
	return EmptyHandError{}
}
func NewFullTableAreaError() FullTableAreaError {
	return FullTableAreaError{}
}