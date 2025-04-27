package errors

import (
	"fmt"
	"strings"
)

type IndexError struct {
	index *int
}
type EmptyError struct{}
type FullError struct{}

func NewIndexError(index *int) IndexError {
	return IndexError{index: index}
}
func NewEmptyError() IndexError {
	return IndexError{}
}
func NewFullError() IndexError {
	return IndexError{}
}

func (err IndexError) Error() string {
	builder := strings.Builder{}
	builder.WriteString("Invalid index")
	if err.index != nil {
		fmt.Fprintf(&builder, ": %d", *err.index)
	}
	return builder.String()
}
func (err EmptyError) Error() string {
	return "Collection is empty"
}
func (err FullError) Error() string {
	return "Collection is full"
}