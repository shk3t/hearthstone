package loop

import "fmt"

type InvalidArgumentsError struct {
	correctUsage string
}

func NewInvalidArgumentsError(correctUsage string) InvalidArgumentsError {
	return InvalidArgumentsError{correctUsage}
}

func (err InvalidArgumentsError) Error() string {
	return fmt.Sprintf("Некорректные аргументы\n%s", err.correctUsage)
}