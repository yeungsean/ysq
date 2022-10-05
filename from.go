package ysq

import (
	"golang.org/x/exp/constraints"
)

// FromElement element list -> Query
func FromElement[T any](source ...T) *Query[T] {
	var t Query[T]
	length := len(source)
	t.Next = func() Iterator[T] {
		idx := 0
		return func() (item T, ok bool) {
			if ok = idx < length; ok {
				item = source[idx]
				idx++
			}
			return
		}
	}
	return &t
}

// FromSlice slice -> Query
func FromSlice[T any](source []T) *Query[T] {
	return FromElement(source...)
}

// FromString string -> Query
func FromString(s string) *Query[rune] {
	tmp := []rune(s)
	return FromElement(tmp...)
}

// FromSequence sequence -> Query
func FromSequence[T constraints.Integer](start, end T, stepE ...int) *Query[T] {
	if start == end {
		panic("end must be greater than start")
	}

	length := end - start + 1
	if length <= 0 {
		panic("end - start + 1 must be greater than 0")
	}

	var step T = 1
	if len(stepE) > 0 {
		step = (T)(stepE[0])
	}
	slice := make([]T, 0, end-start+1)
	for i := start; i <= end; i += step {
		slice = append(slice, i)
	}
	return FromElement(slice...)
}

// FromSequenceChan sequence chan -> Query
func FromSequenceChan[T constraints.Integer](start, end T, stepE ...int) *Query[T] {
	if start < 0 || end < 0 {
		panic("start or end must be greater than or equal 0")
	} else if start == end {
		panic("end must be greater than start")
	}

	var step T = 1
	if len(stepE) > 0 {
		step = (T)(stepE[0])
	}
	ch := make(chan T)
	go func() {
		for i := start; i <= end; i += step {
			ch <- i
		}
		close(ch)
	}()
	return FromChan(ch)
}

// FromChan chan -> Query
func FromChan[T any](c chan T) *Query[T] {
	var t Query[T]
	t.Next = func() Iterator[T] {
		return func() (item T, ok bool) {
			item, ok = <-c
			return
		}
	}
	return &t
}

// FromMap map -> Query
func FromMap[K comparable, V any](m map[K]V) *Query[KeyValuePair[K, V]] {
	var t Query[KeyValuePair[K, V]]
	length := len(m)
	t.Next = func() Iterator[KeyValuePair[K, V]] {
		idx := 0
		keys := make([]K, 0, length)
		for k := range m {
			keys = append(keys, k)
		}
		return func() (item KeyValuePair[K, V], ok bool) {
			if ok = idx < length; ok {
				key := keys[idx]
				item = KeyValuePair[K, V]{
					Key:   key,
					Value: m[key],
				}
				idx++
			}
			return
		}
	}
	return &t
}
