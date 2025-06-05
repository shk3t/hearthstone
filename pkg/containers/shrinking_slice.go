package containers

import (
	errorpkg "hearthstone/pkg/errors"
)

// Shrinking slice.
// Works with pointers and interfaces.
// AVOID direct indexing!
type Shrice[T comparable] []T

func NewShrice[T comparable](size int) Shrice[T] {
	return make([]T, size, size)
}

// Returns number of first non-nil consecutive elements
func (s Shrice[T]) Len() int {
	var null T
	for i, v := range s {
		if v == null {
			return i
		}
	}
	return len(s)
}

func (s Shrice[T]) Cap() int {
	return cap(s)
}

func (s Shrice[T]) Get(idx int) (T, error) {
	var null T
	if idx < 0 || s.Cap() <= idx || s[idx] == null {
		return null, errorpkg.NewIndexError(idx)
	}
	return s[idx], nil
}

func (s Shrice[T]) Insert(idx int, value T) error {
	var null T
	if idx < 0 || s.Cap() <= idx {
		return errorpkg.NewIndexError(idx)
	}
	sLen := s.Len()
	if s.Cap() == sLen {
		return errorpkg.NewFullError()
	}

	if s[idx] != null {
		// Push everything one element to the right
		for i := sLen; i > idx; i-- {
			s[i] = s[i-1]
		}
		s[idx] = value
	} else {
		s[idx] = value
		s.Shrink()
	}

	return nil
}

func (s Shrice[T]) PushBack(values ...T) error {
	sCap := s.Cap()
	sLen := s.Len()
	vLen := len(values)

	if sCap-sLen < vLen {
		return errorpkg.NewNotEnoughSpaceError(sCap-sLen, vLen)
	}

	for i := 0; i < vLen; i++ {
		s[sLen+i] = values[i]
	}
	return nil
}

func (s Shrice[T]) Pop(idx int) (T, error) {
	var null T
	if s[0] == null {
		return null, errorpkg.NewEmptyError()
	} else if idx < 0 || s.Cap() <= idx || s[idx] == null {
		return null, errorpkg.NewIndexError(idx)
	}

	value := s[idx]
	s[idx] = null
	s.Shrink()
	return value, nil
}

func (s Shrice[T]) PopBack() (T, error) {
	var null T
	idx := s.Len() - 1

	if idx < 0 {
		return null, errorpkg.NewEmptyError()
	}

	value := s[idx]
	s[idx] = null
	return value, nil
}

func (s Shrice[T]) Shrink() {  // BUG ???
	var null T
	j := s.Len()
	for i := j + 1; i < s.Cap(); i++ {
		if s[i] != null {
			s[j] = s[i]
			s[i] = null
			j++
		}
	}
}