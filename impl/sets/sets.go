package sets

import "sort"

type SetType interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | string
}

type Set[T SetType] struct {
	elems map[T]struct{}
}

func NewSet[T SetType](elems ...T) *Set[T] {
	s := &Set[T]{
		elems: make(map[T]struct{}, len(elems)),
	}
	s.Add(elems...)
	return s
}

func (s *Set[T]) Len() int {
	return len(s.elems)
}

func (s *Set[T]) IsEmpty() bool {
	return len(s.elems) == 0
}

func (s *Set[T]) Elements() []T {
	elems := make([]T, 0, len(s.elems))
	for elem := range s.elems {
		elems = append(elems, elem)
	}
	sort.Slice(elems, func(i, j int) bool {
		return elems[i] < elems[j]
	})
	return elems
}

func (s *Set[T]) Contains(elem T) bool {
	_, ok := s.elems[elem]
	return ok
}

func (s *Set[T]) Add(elems ...T) {
	for _, elem := range elems {
		s.elems[elem] = struct{}{}
	}
}

func (s *Set[T]) Remove(elems ...T) {
	for _, elem := range elems {
		delete(s.elems, elem)
	}
}

func (s *Set[T]) Clear() {
	s.elems = make(map[T]struct{})
}

func (s *Set[T]) Union(b *Set[T]) *Set[T] {
	r := NewSet(s.Elements()...)
	r.Add(b.Elements()...)
	return r
}

func (s *Set[T]) Intersect(b *Set[T]) *Set[T] {
	r := NewSet[T]()
	for elem := range s.elems {
		if b.Contains(elem) {
			r.Add(elem)
		}
	}
	return r
}

func (s *Set[T]) Difference(b *Set[T]) *Set[T] {
	r := NewSet(s.Elements()...)
	for elem := range b.elems {
		delete(r.elems, elem)
	}
	return r
}
