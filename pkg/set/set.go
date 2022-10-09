// Package set ...
package set

import (
	"github.com/yeungsean/ysq/pkg/delegate"
)

// Set is a set of values
type Set[T any, TKey comparable] struct {
	m           map[TKey]T
	keySelector delegate.Func1[T, TKey]
}

// MakeSet returns a set of some element type.
func MakeSet[T any, TKey comparable](keySelector delegate.Func1[T, TKey]) *Set[T, TKey] {
	return &Set[T, TKey]{
		m:           make(map[TKey]T),
		keySelector: keySelector,
	}
}

// Add adds v to the set s.
func (s *Set[T, TKey]) Add(vs ...T) {
	for _, v := range vs {
		s.m[s.keySelector(v)] = v
	}
}

// Delete removes v from the set s.
func (s *Set[T, TKey]) Delete(v T) {
	delete(s.m, s.keySelector(v))
}

// Contains reports whether v is in s.
func (s *Set[T, TKey]) Contains(v T) bool {
	_, ok := s.m[s.keySelector(v)]
	return ok
}

// Len reports the number of elements in s.
func (s *Set[T, TKey]) Len() int {
	return len(s.m)
}

// Iterate invokes f on each element of s.
// It's OK for f to call the Delete method.
func (s *Set[T, TKey]) Iterate(f delegate.Action1[T]) {
	for _, v := range s.m {
		f(v)
	}
}

// ToSlice ...
func (s *Set[T, TKey]) ToSlice() []T {
	slice := make([]T, 0, s.Len())
	for _, v := range s.m {
		slice = append(slice, v)
	}
	return slice
}
