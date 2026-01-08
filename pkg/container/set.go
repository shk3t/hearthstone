package container

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](values []T) Set[T] {
	set := make(map[T]struct{}, len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}
	return set
}

func (s Set[T]) Has(value T) bool {
	_, ok := s[value]
	return ok
}
