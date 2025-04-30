package helpers

import (
	"fmt"
	"strings"
)

func JoinErrors(errs []error, sep string) string {
	builder := strings.Builder{}
	for _, err := range errs {
		builder.WriteString(err.Error())
		builder.WriteString(sep)
	}
	return strings.TrimSuffix(builder.String(), sep)
}

func UnexpectedError(err error) string {
	return fmt.Sprintf("Unexpected error: %v", err)
}