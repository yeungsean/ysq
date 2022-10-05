package ysq

import (
	"github.com/yeungsean/ysq/pkg/delegate"
)

// Skip 跳过序列中指定数量的元素，然后返回剩余的元素。
func (q *Query[T]) Skip(cnt int) *Query[T] {
	return &Query[T]{
		Next: func() Iterator[T] {
			next := q.Next()
			n := cnt
			return func() (item T, ok bool) {
				for ; n > 0; n-- {
					if item, ok = next(); !ok {
						return
					}
				}

				return next()
			}
		},
	}
}

// SkipWhile 如果指定的条件为 true，则跳过序列中的元素，然后返回剩余的元素。
func (q *Query[T]) SkipWhile(predicate delegate.FuncTBool[T]) *Query[T] {
	return &Query[T]{
		Next: func() Iterator[T] {
			next, loop := q.Next(), true
			return func() (item T, ok bool) {
				for loop {
					if item, ok = next(); !ok {
						return
					} else if !predicate(item) {
						continue
					}
					loop = false
					return
				}

				return next()
			}
		},
	}
}
