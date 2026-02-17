package loop

import errpkg "hearthstone/pkg/errors"

// All game errors have to embed this structure.
// In this case they become instances of `error` type.
//
// `BaseError` `Error()` is not supposed to be used.
// Use your to string conversion function instead.
type BaseError struct{}

func (err BaseError) Error() string {
	panic(errpkg.NewUnusableFeatureError())
}

type NeedActionError struct{
	BaseError
}

func NewNeedActionError() NeedActionError {
	return NeedActionError{}
}
