package collections

import "errors"

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
		return errors.New("Invalid index")
	}
	length := s.Len()
	if s.Cap() == length {
		return errors.New("Shrice is full")
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

func (s Shrice[T]) Pop(idx int) (T, error) {
	var null T
	if idx < 0 || s.Cap() <= idx || s[idx] == null {
		return null, errors.New("Invalid index")
	}

	value := s[idx]
	s[idx] = null
	s.shrink()
	return value, nil
}

func (s Shrice[T]) shrink() {
	var null T
	j := s.Len() - 1
	for i := j + 1; i < s.Cap(); i++ {
		if s[i] != null {
			s[j] = s[i]
			s[i] = null
			j++
		}
	}
}