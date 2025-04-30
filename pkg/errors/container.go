package errors

import (
	"fmt"
)

type IndexError struct {
	index int
}
type EmptyError struct{}
type FullError struct{}

func NewIndexError(index int) IndexError {
	return IndexError{index: index}
}
func NewEmptyError() EmptyError {
	return EmptyError{}
}
func NewFullError() FullError {
	return FullError{}
}

func (err IndexError) Error() string {
	return fmt.Sprintf("Invalid index: %d", err.index)
}
func (err EmptyError) Error() string {
	return "Collection is empty"
}
func (err FullError) Error() string {
	return "Collection is full"
}