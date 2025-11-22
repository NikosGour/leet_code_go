package internal

type Set[T comparable] struct {
	_map map[T]bool
}

func (s *Set[T]) append(val T) {
	if !s.contains(val) {
		s._map[val] = true
	}
}
func (s *Set[T]) contains(val T) bool {
	_, exist := s._map[val]
	return exist
}

func (s *Set[T]) size() int {
	return len(s._map)
}

func NewSet[T comparable]() Set[T] {
	_map := make(map[T]bool)
	return Set[T]{_map: _map}
}
