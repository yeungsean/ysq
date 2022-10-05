package ysq

import (
	"github.com/yeungsean/ysq/pkg/delegate"
)

// Filter alias Where
func (q *Query[T]) Filter(predicate delegate.FuncTBool[T]) *Query[T] {
	return q.Where(predicate)
}

// Where 基于谓词筛选值序列
func (q *Query[T]) Where(predicate delegate.FuncTBool[T]) *Query[T] {
	return &Query[T]{
		Next: func() Iterator[T] {
			next := q.Next()
			return func() (item T, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if predicate(item) {
						return
					}
				}
				return
			}
		},
	}
}

// FilterN alias WhereN
func (q *Query[T]) FilterN(predicate delegate.FuncTIntBool[T]) *Query[T] {
	return q.WhereN(predicate)
}

// WhereN 基于谓词筛选值序列。 将在谓词函数的逻辑中使用每个元素的索引
func (q *Query[T]) WhereN(predicate delegate.FuncTIntBool[T]) *Query[T] {
	return &Query[T]{
		Next: func() Iterator[T] {
			next := q.Next()
			idx := 0
			return func() (item T, ok bool) {
				for item, ok = next(); ok; item, ok = next() {
					if predicate(item, idx) {
						idx++
						return
					}
					idx++
				}
				return
			}
		},
	}
}
