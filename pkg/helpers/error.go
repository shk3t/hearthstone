package helpers

import (
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

func FirstError(errs []error) error {
	for _, err := range errs{
		if err != nil {
			return err
		}
	}
	return nil
}