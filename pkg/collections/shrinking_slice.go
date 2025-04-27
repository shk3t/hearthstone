package collections

import (
	errorpkg "hearthstone/pkg/errors"
)

// Shrinking slice
// Works with pointers and interfaces
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

func (s Shrice[T]) Insert(idx int, value T) error {
	var null T
	if idx < 0 || s.Cap() <= idx {
		return errorpkg.NewIndexError(&idx)
	}
	length := s.Len()
	if s.Cap() == length {
		return errorpkg.NewFullError()
	}

	if s[idx] != null {
		// Push everything one element to the right
		for i := length; i > idx; i-- {
			s[i] = s[i-1]
		}
		s[idx] = value
	} else {
		s[idx] = value
		s.shrink()
	}

	return nil
}

// func (s Shrice[T]) PushBack(value T) error {
// 	// TODO: add last
// }

func (s Shrice[T]) Pop(idx int) (T, error) {
	var null T
	// TODO: empty
	if idx < 0 || s.Cap() <= idx || s[idx] == null {
		return null, errorpkg.NewIndexError(&idx)
	}

	value := s[idx]
	s[idx] = null
	s.shrink()
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

func (s Shrice[T]) shrink() {
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