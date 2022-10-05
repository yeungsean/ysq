package ysq

import (
	"github.com/yeungsean/ysq/pkg/delegate"
)

// Take 包含输入序列开头的指定数量的元素
func (q *Query[T]) Take(cnt uint) *Query[T] {
	if cnt <= 0 {
		panic("cnt must be greater than 0")
	}
	return &Query[T]{
		Next: func() Iterator[T] {
			next, idx := q.Next(), cnt
			return func() (item T, ok bool) {
				if idx <= 0 {
					return
				}

				idx--
				return next()
			}
		},
	}
}

// TakeWhile 只要指定的条件为 true，就会返回序列的元素
func (q *Query[T]) TakeWhile(predicate delegate.FuncTBool[T]) *Query[T] {
	return &Query[T]{
		Next: func() Iterator[T] {
			next := q.Next()
			return func() (item T, ok bool) {
				if item, ok = next(); !ok {
					return
				} else if predicate(item) {
					return
				}

				var zero T
				return zero, false
			}
		},
	}
}

// TakeWhileN 只要指定的条件为 true，就会返回序列的元素。将在谓词函数的逻辑中使用元素的索引
func (q *Query[T]) TakeWhileN(predicate delegate.FuncTIntBool[T]) *Query[T] {
	return &Query[T]{
		Next: func() Iterator[T] {
			next := q.Next()
			idx := -1
			return func() (item T, ok bool) {
				idx++
				if item, ok = next(); !ok {
					return
				} else if predicate(item, idx) {
					return
				}

				var zero T
				return zero, false
			}
		},
	}
}
