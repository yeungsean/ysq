package ysq

import "github.com/yeungsean/ysq/pkg/delegate"

// Reduce reduces a []T1 to a single value using a reduction function.
func Reduce[T, TResult any](q *Query[T], initializer TResult, f delegate.Func2[TResult, T, TResult]) TResult {
	r := initializer
	next := q.Next()
	for item, ok := next(); ok; item, ok = next() {
		r = f(r, item)
	}
	return r
}

// Reduce reduces a []T1 to a single value using a reduction function.
func (q *Query[T]) Reduce(initializer T, f delegate.Func2[T, T, T]) T {
	return Reduce(q, initializer, f)
}
