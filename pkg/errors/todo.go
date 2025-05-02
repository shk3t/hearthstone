package errors

import (
	"fmt"
	"hearthstone/pkg/helpers"
)

type NotImplementedError struct {
	feature string
}
type UnexpectedError struct {
	baseErr error
}
type UnusableFeatureError struct{}

func NewNotImplementedError(feature string) NotImplementedError {
	return NotImplementedError{feature}
}
func NewUnexpectedError(baseErr error) UnexpectedError {
	return UnexpectedError{baseErr}
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
func (err UnexpectedError) Error() string {
	return fmt.Sprintf("Unexpected error: %s", err.baseErr)
}
func (err UnusableFeatureError) Error() string {
	return "This feature is not intended to be used"
}