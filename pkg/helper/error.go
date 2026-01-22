package helper

import (
	"fmt"
	"strings"
)

func JoinErrors(errs []error, sep string) string {
	builder := strings.Builder{}
	for _, err := range errs {
		fmt.Fprint(&builder, err.Error())
		fmt.Fprint(&builder, sep)
	}
	return strings.TrimSuffix(builder.String(), sep)
}

func FirstError(errs []error) error {
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return nil
}