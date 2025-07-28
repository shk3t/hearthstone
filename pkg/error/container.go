package error

import (
	"fmt"
)

type IndexError struct {
	index int
}
type EmptyError struct{}
type FullError struct{}
type NotEnoughSpaceError struct {
	available int
	required  int
}

func NewIndexError(index int) IndexError {
	return IndexError{index: index}
}
func NewEmptyError() EmptyError {
	return EmptyError{}
}
func NewFullError() FullError {
	return FullError{}
}
func NewNotEnoughSpaceError(available, required int) NotEnoughSpaceError {
	return NotEnoughSpaceError{available, required}
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
func (err NotEnoughSpaceError) Error() string {
	return fmt.Sprintf("Not enough space. Available: %d, required: %d", err.available, err.required)
}