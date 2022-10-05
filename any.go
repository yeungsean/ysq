package ysq

import (
	"github.com/yeungsean/ysq/pkg/delegate"
)

// Contains alias In
func (q *Query[T]) Contains(predicate delegate.FuncTBool[T]) bool {
	return q.In(predicate)
}

// In 确定序列是否包含任何元素
func (q *Query[T]) In(predicate delegate.FuncTBool[T]) bool {
	next := q.Next()
	for item, ok := next(); ok; item, ok = next() {
		if predicate(item) {
			return true
		}
	}
	return false
}

// IsEmpty 序列是否为空
func (q *Query[T]) IsEmpty() bool {
	_, ok := q.Next()()
	return !ok
}

// Any 序列是否不为空
func (q *Query[T]) Any() bool {
	return !q.IsEmpty()
}

// All 确定序列中的所有元素是否都满足条件
func (q *Query[T]) All(predicate delegate.FuncTBool[T]) bool {
	next := q.Next()
	for item, ok := next(); ok; item, ok = next() {
		if !predicate(item) {
			return false
		}
	}
	return true
}
