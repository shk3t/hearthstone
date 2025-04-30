package errors

import (
	"fmt"
	"hearthstone/pkg/helpers"
)

type NotImplementedError struct {
	feature string
}
type UnusableFeatureError struct{}

func NewNotImplementedError(feature string) NotImplementedError {
	return NotImplementedError{feature}
}
func NewUnusableFeatureError() UnusableFeatureError {
	return UnusableFeatureError{}
}

func (err NotImplementedError) Error() string {
	return fmt.Sprintf(
		"%s %s not implemented",
		helpers.Capitalize(err.feature),
		helpers.BeForm(err.feature),
	)
}

func (err UnusableFeatureError) Error() string {
	return "This feature is not intended to be used"
}