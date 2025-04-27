package errors

import (
	"fmt"
	"strings"
)

type NotImplementedError struct {
	feature string
}

func (err NotImplementedError) Error() string {
	return fmt.Sprintf("%s is not implemented", strings.ToTitle(err.feature))
}

func NewNotImplementedError(feature string) NotImplementedError {
	return NotImplementedError{feature}
}