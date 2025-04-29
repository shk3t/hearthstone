package errors

import (
	"fmt"
	"hearthstone/pkg/helpers"
)

type NotImplementedError struct {
	feature string
}

func (err NotImplementedError) Error() string {
	return fmt.Sprintf(
		"%s %s not implemented",
		helpers.Capitalize(err.feature),
		helpers.BeForm(err.feature),
	)
}

func NewNotImplementedError(feature string) NotImplementedError {
	return NotImplementedError{feature}
}